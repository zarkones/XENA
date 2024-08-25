package ctrl

import (
	attacksRepo "c2/repos/attacks"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOngoingAttacks(c *gin.Context) {
	attacks, err := attacksRepo.GetMultiple()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if attacks == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, attacks)
}

func AttackTarget(c *gin.Context) {
	// TODO
}
