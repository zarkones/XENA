package ctrl

import (
	"c2/models"
	messagesRepo "c2/repos/messages"
	pipelinesRepo "c2/repos/pipelines"
	"common/slices"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	xenaC2 "github.com/zarkones/xena-client"

	"github.com/gin-gonic/gin"
)

func GetPipelines(c *gin.Context) {
	pipelines, err := pipelinesRepo.GetMultiple()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(pipelines) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, pipelines)
}

func DeletePipeline(c *gin.Context) {
	pipelineID := c.Param("pipelineID")

	if err := pipelinesRepo.Delete(pipelineID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
}

func UpsertPipeline(c *gin.Context) {
	var pipe xenaC2.Pipeline

	if err := c.BindJSON(&pipe); err != nil {
		fmt.Println("failed to unserialize pipeline:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var settings xenaC2.PipelineSettings

	if err := json.Unmarshal([]byte(pipe.Settings), &settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	// Deduplicate .LinkedTo, as we cannot have one node being linked multiple times to another.
	for id := range settings.Steps {
		settings.Steps[id] = xenaC2.PipelineStep{
			ID:       settings.Steps[id].ID,
			Name:     settings.Steps[id].Name,
			Position: settings.Steps[id].Position,
			Tool:     settings.Steps[id].Tool,
			LinkedTo: slices.Deduplicate(settings.Steps[id].LinkedTo),
		}
	}

	if err := pipelinesRepo.Upsert(&models.Pipeline{
		ID:       pipe.ID,
		Name:     pipe.Name,
		Desc:     pipe.Desc,
		Category: pipe.Category,
		Settings: pipe.Settings,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
}

func ExecPipeline(c *gin.Context) {
	var execCtx xenaC2.ExecPipelineReqCtx

	if err := c.BindJSON(&execCtx); err != nil {
		fmt.Println("failed to unserialize pipeline execution contract:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	pipeline, err := pipelinesRepo.Get(execCtx.PipelineID)
	if err != nil {
		fmt.Println("failed to get pipeline:", execCtx.PipelineID, err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	jsonPipeline, err := json.Marshal(pipeline)
	if err != nil {
		fmt.Println("failed to marhsal pipeline:", execCtx.PipelineID, err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	pipelineExecution := models.PipelineRun{
		ID:         uuid.NewString(),
		PipelineID: pipeline.ID,
		ExecutedAt: time.Now().Unix(),
	}

	if err := pipelinesRepo.RunUpsert(&pipelineExecution); err != nil {
		fmt.Println("failed to mark pipeline execution:", pipeline.ID, err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	msg := models.Message{
		ID:                  uuid.NewString(),
		PipelineExecutionID: pipelineExecution.ID,
		AgentID:             execCtx.AgentIDs[0], // Supporting only single agent execution mode atm.
		FriendlyTitle:       "Executing Pipeline " + pipeline.Name,
		Request:             string(jsonPipeline),
	}

	if err := messagesRepo.Insert(&msg); err != nil {
		fmt.Println("failed to insert message of pipeline:", pipeline.ID, err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}

func SetPipelineSettings(c *gin.Context) {
	var pipe xenaC2.SetPipelineSettingsReqCtx

	if err := c.BindJSON(&pipe); err != nil {
		fmt.Println("failed to unserialize pipeline:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	serializedSettings, err := json.Marshal(pipe.PipelineSettings)
	if err != nil {
		fmt.Println("failed to serialize pipeline's settings:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := pipelinesRepo.SetSettings(pipe.PipelineID, string(serializedSettings)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
}

func GetPipelineExecutions(c *gin.Context) {
	pipelineID := c.Param("pipelineID")

	runs, err := pipelinesRepo.GetRuns(pipelineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	if len(runs) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, runs)
}
