package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pp "github.com/husamettinarabaci/go-resthell/core/application/presentation/port"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	dto "github.com/husamettinarabaci/go-resthell/pkg/presentation/dto"
)

type RestAPI struct {
	engine         *gin.Engine
	commandHandler pp.CommandPort
	queryHandler   pp.QueryPort
}

func NewRestAPI(qh pp.QueryPort, ch pp.CommandPort) *RestAPI {
	api := &RestAPI{
		commandHandler: ch,
		queryHandler:   qh,
	}
	api.engine = gin.Default()
	api.engine.POST("/api/cmd", api.PostCommand)
	return api
}

func (api *RestAPI) Serve(debug bool, port string) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	api.engine.Run(":" + port)
}

func (api *RestAPI) PostCommand(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd, ok := json["command"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": mo.ErrInvalidInput.Error()})
		return
	}
	var commandRequestDto dto.CommandRequest
	commandRequestDto.Command = cmd
	if commandRequestDto.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": mo.ErrInvalidInput.Error()})
		return
	}
	isCommandExists, _ := api.queryHandler.IsCommandExists(c, commandRequestDto.ToCommandRequestEntity())
	if !isCommandExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": mo.ErrCommandIsNotShellCommand.Error()})
		return
	}
	executeCommandResponse, err := api.commandHandler.ExecuteCommand(c, commandRequestDto.ToCommandRequestEntity())
	commandResponseDto := dto.FromResponseObject(executeCommandResponse.Response)
	if err != nil {
		commandResponseDto.Err = append(commandResponseDto.Err, err.Error())
		c.JSON(http.StatusBadRequest, commandResponseDto)
		return
	}
	c.JSON(http.StatusOK, commandResponseDto)
}
