package application

import (
	"context"

	as "github.com/husamettinarabaci/go-resthell/core/application/service"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
)

type CommandAdapter struct {
	service as.Service
}

func NewCommandAdapter(s as.Service) CommandAdapter {
	return CommandAdapter{
		service: s,
	}
}

func (a CommandAdapter) ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (me.CommandResponse, error) {
	return a.service.ExecuteCommand(ctx, commandRequest)
}
