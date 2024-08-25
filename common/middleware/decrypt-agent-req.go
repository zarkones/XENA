package middleware

import (
	"bytes"
	"c2/core"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	cry "github.com/zarkones/xena-crypto"
)

func DecryptAgentReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("DecryptAgentReq: io.ReadAll():", err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		decryptedBody, err := cry.SecureUnwrap(core.PrivateKey, string(reqBody))
		if err != nil {
			fmt.Println("DecryptAgentReq: cry.SecureUnwrap():", err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decryptedBody)))

		c.Next()
	}
}
