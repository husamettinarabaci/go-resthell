package application

import (
	"context"

	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
)

type CommandPort interface {
	ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (me.CommandResponse, error)
}
