package ctrl

import (
	"c2/core/env"
	"c2/core/pubsub"
	"c2/models"
	agentsRepo "c2/repos/agents"
	messagesRepo "c2/repos/messages"
	pipelinesRepo "c2/repos/pipelines"
	targetsRepo "c2/repos/targets"
	"common/debug"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	xenaC2 "github.com/zarkones/xena-client"
	cry "github.com/zarkones/xena-crypto"
	"gorm.io/gorm"
)

func MessagesSubscribe(c *gin.Context) {
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	agent, err := agentsRepo.Get(c.Param("agentID"))
	if err != nil {
		debug.Println("ctrl.MessagesSubscribe: agentsRepo.Get():", err)
		c.Status(http.StatusNotFound)
		return
	}

	agentPubKey, err := cry.ImportPubKeyPEM([]byte(agent.PubKeyPEM))
	if err != nil {
		debug.Println("ctrl.MessagesSubscribe: cry.ImportPubKeyPEM():", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	reqCtx, cancelReqCtx := context.WithCancel(c.Request.Context())
	defer func() {
		cancelReqCtx()
		close(pubsub.Messages[agent.ID])
		delete(pubsub.Messages, agent.ID)
	}()

	go func() {
		if pubsub.Messages[agent.ID] == nil {
			pubsub.Messages[agent.ID] = make(chan xenaC2.Message)
		}

		sendMessage := func(msg xenaC2.Message) {
			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("ctrl.MessagesSubscribe: json.Marshal a message failed:", err)
				return
			}

			encryptedMsg, err := cry.EncryptRSAOAEPEncodeHex(*agentPubKey, string(jsonMsg))
			if err != nil {
				fmt.Println("ctrl.MessagesSubscribe: cry.EncryptRSAOAEPEncodeHex():", err)
				return
			}

			jsonMsg = append([]byte(encryptedMsg), []byte(xenaC2.MESSAGE_STREAM_SEPARATOR+"\n")...)

			if _, err := c.Writer.Write(jsonMsg); err != nil {
				fmt.Println("ctrl.MessagesSubscribe: c.Writer.Write failed to send a message:", err)
				return
			}

			flusher.Flush()
		}

		// TODO: Maybe handle the error.
		pendingMessages, _ := messagesRepo.GetMultipleForAgent(agent.ID)
		for _, pendingMessage := range pendingMessages {
			sendMessage(xenaC2.Message(pendingMessage))
		}

		for msg := range pubsub.Messages[agent.ID] {
			sendMessage(msg)
		}
	}()

	<-reqCtx.Done()
}

func MessageGet(c *gin.Context) {
	messageID := c.Request.URL.Query().Get("messageID")
	agentID := c.Request.URL.Query().Get("agentID")
	request := c.Request.URL.Query().Get("request")

	message, err := func() (models.Message, error) {
		if messageID != "" {
			return messagesRepo.Get(messageID)
		}
		if agentID != "" && request != "" {
			return messagesRepo.GetByRequest(agentID, request)
		}
		return models.Message{}, errors.New("bad user input")
	}()
	if err != nil {
		fmt.Println("ctrl.MessageGetByReq: failed to retrieve message:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	jsonMessage, err := json.Marshal(&message)
	if err != nil {
		fmt.Println("ctrl.MessageGetByReq: json.Marshal():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.Write([]byte(jsonMessage))
}

func MessagesGetMultiple(c *gin.Context) {
	agentID := c.Param("agentID")

	msgHandler := messagesRepo.GetMultipleForAgent
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && authHeader == env.AUTH_TOKEN {
		msgHandler = messagesRepo.GetMultiple
	}

	messages, err := msgHandler(agentID)
	if err != nil {
		fmt.Println("ctrl.MessagesGetMultiple: failed to retrieve messages:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(messages) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	if authHeader != "" {
		c.JSON(http.StatusOK, messages)
		return
	}

	agent, err := agentsRepo.Get(agentID)
	if err != nil {
		fmt.Println("ctrl.MessagesGetMultiple: agentsRepo.Get():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	agentPubKey, err := cry.ImportPubKeyPEM([]byte(agent.PubKeyPEM))
	if err != nil {
		fmt.Println("ctrl.MessagesGetMultiple: cry.ImportPubKeyPEM():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	jsonMessages, err := json.Marshal(&messages)
	if err != nil {
		fmt.Println("ctrl.MessagesGetMultiple: json.Marshal():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	encryptedResp, err := cry.SecureWrap(agentPubKey, string(jsonMessages))
	if err != nil {
		fmt.Println("ctrl.MessagesGetMultiple: cry.SecureWrap():", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.Write([]byte(encryptedResp))
}

func MessagesInsert(c *gin.Context) {
	var newMsg models.Message

	if err := c.BindJSON(&newMsg); err != nil {
		fmt.Println("failed to unserialize message:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if newMsg.Response != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot insert message with response"})
		return
	}

	if err := messagesRepo.Insert(&newMsg); err != nil {
		fmt.Println("failed to insert message:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if _, ok := pubsub.Messages[newMsg.AgentID]; ok {
		pubsub.Messages[newMsg.AgentID] <- xenaC2.Message(newMsg)
	}

	c.JSON(http.StatusCreated, gin.H{"id": newMsg.ID})
}

func MessagesAddResponse(c *gin.Context) {
	var msgResp xenaC2.AgentMsgRespCtx
	if err := c.BindJSON(&msgResp); err != nil {
		fmt.Println("failed to unserialize message response:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var maybePipelinePayload xenaC2.Pipeline
	if err := json.Unmarshal([]byte(msgResp.Response), &maybePipelinePayload); err == nil {
		go func() {
			if processPipeline(msgResp.PipelineExecutionID, msgResp.Response); err != nil {
				fmt.Println("Pipeline Processing Failed:", maybePipelinePayload.ID, err)
			}
		}()
	}

	if err := messagesRepo.UpdateResponse(msgResp.MessageID, msgResp.Response); err != nil {
		fmt.Println("failed to update message response:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func processPipeline(pipelineRunID, serializedPipeline string) error {
	if pipelineRunID == "" {
		return nil
	}

	run, err := pipelinesRepo.GetRun(pipelineRunID)
	if err != nil {
		return err
	}

	run.FinishedPipeline = serializedPipeline
	run.Finished()

	if err := pipelinesRepo.RunUpsert(&run); err != nil {
		return err
	}

	var finishedPipeline xenaC2.Pipeline
	if err := json.Unmarshal([]byte(run.FinishedPipeline), &finishedPipeline); err != nil {
		return err
	}
	var settings xenaC2.PipelineSettings
	if err := json.Unmarshal([]byte(finishedPipeline.Settings), &settings); err != nil {
		return err
	}

	for _, step := range settings.Steps {
		switch step.Tool.ID {
		case "SUBDOMAIN_ENUM_TOP100", "SUBDOMAIN_ENUM_TOP500", "SUBDOMAIN_ENUM_TOP1K", "SUBDOMAIN_ENUM_TOP10K", "SUBDOMAIN_ENUM_PASSIVE":
			rawDomains, ok := step.Tool.Outputs["stdout"]
			if !ok {
				continue
			}
			domains := strings.Split(rawDomains.Value, "\n")
			for _, domain := range domains {
				if domain == "" {
					continue
				}
				if _, err := targetsRepo.GetByValue(domain); err == gorm.ErrRecordNotFound {
					if err := targetsRepo.Upsert(&models.Target{
						Value: domain,
						Type:  "DOMAIN",
					}); err != nil {
						fmt.Println("Failed to insert target:", err)
						continue
					}
				}
			}
		}
	}

	return nil
}
