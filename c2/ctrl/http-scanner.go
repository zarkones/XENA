package ctrl

import (
	"c2/models"
	"c2/repos/httpScansRepo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_HTTP_SCANS_PER_PAGE = 28
)

func GetHttpScans(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	offset := DEFAULT_HTTP_SCANS_PER_PAGE * page

	scans, err := httpScansRepo.GetMultiple(DEFAULT_HTTP_SCANS_PER_PAGE, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(scans) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, scans)
}

func InsertHttpScan(c *gin.Context) {
	var scan models.HttpScan

	if err := c.BindJSON(&scan); err != nil {
		fmt.Println("failed to unserialize http scan:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := httpScansRepo.Insert(scan.ReqID, scan.AgentIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}
