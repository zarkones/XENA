package ctrl

import (
	"c2/core"
	"net/http"

	"github.com/gin-gonic/gin"
	cry "github.com/zarkones/xena-crypto"
)

func GetC2PublicKey(c *gin.Context) {
	pubKey, err := cry.PubKeyToPEM(&core.PrivateKey.PublicKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.Writer.Write([]byte(pubKey))
}
