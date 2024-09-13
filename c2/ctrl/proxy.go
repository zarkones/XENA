package ctrl

import (
	"c2/core"
	"c2/core/proxy"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PROXIED_REQ_PER_PAGE = 100
)

func GetProxiedRequests(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	offset := DEFAULT_PROXIED_REQ_PER_PAGE * page

	if len(proxy.RequestsStream) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	if len(proxy.RequestsStream) < offset {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	if len(proxy.RequestsStream) < offset+DEFAULT_PROXIED_REQ_PER_PAGE {
		c.JSON(http.StatusOK, proxy.RequestsStream[offset:len(proxy.RequestsStream)-1])
	}

	c.JSON(http.StatusOK, proxy.RequestsStream[offset:offset+DEFAULT_PROXIED_REQ_PER_PAGE])
}

func GetCertificate(c *gin.Context) {
	cert, err := os.ReadFile(core.CertPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.Writer.Header().Add("Content-Disposition", "attachment; filename=\"xena.cert\"")

	c.Writer.Write(cert)
}
