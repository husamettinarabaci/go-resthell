package application

import (
	"context"

	ip "github.com/husamettinarabaci/go-resthell/core/application/infrastructure/port"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	ds "github.com/husamettinarabaci/go-resthell/core/domain/service"
)

type Service struct {
	domainService ds.Service
	shell         ip.ShellPort
}

func NewService(domainService ds.Service, shellPort ip.ShellPort) Service {
	return Service{
		domainService: domainService,
		shell:         shellPort,
	}
}

func (a Service) ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (me.CommandResponse, error) {
	if err := a.domainService.IsCommandRequestEntityValid(commandRequest); err != nil {
		return me.CommandResponse{
			Id: commandRequest.Id,
		}, err
	}
	response, err := a.shell.ExecuteCommand(ctx, commandRequest)
	commandResponse := me.NewCommandResponse(commandRequest.Id, response)
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
