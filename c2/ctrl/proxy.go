package ctrl

import (
	"c2/core"
	proxyRepo "c2/repos/proxy"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PROXIED_REQ_PER_PAGE = 28
)

func GetProxiedRequests(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	offset := DEFAULT_PROXIED_REQ_PER_PAGE * page

	requests, err := proxyRepo.GetMultiple(offset, DEFAULT_PROXIED_REQ_PER_PAGE)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(requests) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetCertificate(c *gin.Context) {
	cert, err := os.ReadFile(core.CertPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.Writer.Header().Add("Content-Disposition", "attachment; filename=\"xena-cert.der\"")

	c.Writer.Write(cert)
}
