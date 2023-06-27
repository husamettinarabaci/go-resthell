package application

import (
	"context"

	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
)

type ShellPort interface {
	IsCommandExists(ctx context.Context, commandRequest me.CommandRequest) (bool, error)
	ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (mo.Response, error)
}
