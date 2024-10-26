package middleware

import (
	"c2/core/env"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OperatorAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != env.AUTH_TOKEN {
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
