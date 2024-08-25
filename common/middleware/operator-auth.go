package middleware

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func OperatorAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != os.Getenv("AUTH_TOKEN") {
			rnd := rand.Intn(10)
			time.Sleep(time.Second * time.Duration(20+rnd))
			c.Writer.WriteHeader(http.StatusServiceUnavailable)

			c.Writer.Header().Del("Access-Control-Allow-Credentials")
			c.Writer.Header().Del("Access-Control-Allow-Headers")
			c.Writer.Header().Del("Access-Control-Allow-Methods")
			c.Writer.Header().Del("Access-Control-Allow-Origin")
			c.Writer.Header().Del("Date")
			c.Abort()
		}

		c.Next()
	}
}
