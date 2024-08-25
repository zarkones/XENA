package ctrl

import (
	"c2/models"
	agentsRepo "c2/repos/agents"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	cry "github.com/zarkones/xena-crypto"
)

func AgentsGetMultiple(c *gin.Context) {
	agents, err := agentsRepo.GetMultiple()
	if err != nil {
		fmt.Println("failed to get agents:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(agents) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, agents)
}

func AgentInsert(c *gin.Context) {
	var newAgent models.Agent

	if err := c.ShouldBindJSON(&newAgent); err != nil {
		fmt.Println("AgentInsert: failed to unserialize agent:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	newAgentPubKey, err := cry.ImportPubKeyPEM([]byte(newAgent.PubKeyPEM))
	if err != nil {
		fmt.Println("AgentInsert: failed parse new agent's public key:", err, hex.EncodeToString([]byte(newAgent.PubKeyPEM)))
		c.JSON(http.StatusForbidden, nil)
		return
	}

	newAgent.IpAddress = c.RemoteIP()

	if err := agentsRepo.Insert(&newAgent); err != nil {
		fmt.Println("AgentInsert: failed to insert agent:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	jsonResp, err := json.Marshal(gin.H{"id": newAgent.ID})
	if err != nil {
		fmt.Println("AgentInsert: failed to marshal JSON response:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	encryptedResp, err := cry.SecureWrap(newAgentPubKey, string(jsonResp))
	if err != nil {
		fmt.Println("AgentInsert: cry.SecureWrap():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.Write([]byte(encryptedResp))
}
