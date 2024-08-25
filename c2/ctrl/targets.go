package ctrl

import (
	"c2/models"
	targetsRepo "c2/repos/targets"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	xenaC2 "github.com/zarkones/xena-client"
)

func GetTargets(c *gin.Context) {
	targets, err := targetsRepo.GetMultiple()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(targets) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, targets)
}

func UpsertTarget(c *gin.Context) {
	var maybeTarget xenaC2.Target

	if err := c.BindJSON(&maybeTarget); err != nil {
		fmt.Println("failed to unserialize target:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	target := &models.Target{
		Type:  maybeTarget.Type,
		Name:  maybeTarget.Name,
		Value: maybeTarget.Value,
	}

	if err := target.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err": err})
		return
	}

	if err := targetsRepo.Upsert(target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
}

func RemoveTarget(c *gin.Context) {
	targetID := c.Param("targetID")

	if err := targetsRepo.Delete(targetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
}
