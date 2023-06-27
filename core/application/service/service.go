package application

import (
	"context"

	ip "github.com/husamettinarabaci/go-resthell/core/application/infrastructure/port"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	mi "github.com/husamettinarabaci/go-resthell/core/domain/model/interface"
	ds "github.com/husamettinarabaci/go-resthell/core/domain/service"
)

type Service struct {
	domainService ds.Service
	shell         ip.ShellPort
	logger        ip.LogPort
}

func NewService(domainService ds.Service, shellPort ip.ShellPort, logger ip.LogPort) Service {
	return Service{
		domainService: domainService,
		shell:         shellPort,
		logger:        logger,
	}
}

func (a Service) Log(ctx context.Context, operationName string, logData mi.Loggable) {
	a.logger.Log(ctx, operationName, logData)
}

func (a Service) ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (me.CommandResponse, error) {
	operationName := "ExecuteCommand"
	a.Log(ctx, operationName, commandRequest)
	if err := a.domainService.IsCommandRequestEntityValid(commandRequest); err != nil {
		return me.CommandResponse{
			Id: commandRequest.Id,
		}, err
	}
	response, err := a.shell.ExecuteCommand(ctx, commandRequest)
	commandResponse := me.NewCommandResponse(commandRequest.Id, response)
	a.Log(ctx, operationName, commandResponse)
	return commandResponse, err
}

func (a Service) IsCommandExists(ctx context.Context, commandRequest me.CommandRequest) (bool, error) {
	if err := a.domainService.IsCommandRequestEntityValid(commandRequest); err != nil {
		return false, err
	}
	isCommandExists, err := a.shell.IsCommandExists(ctx, commandRequest)
	if isCommandExists {
		return true, nil
	}
	return isCommandExists, err
}
