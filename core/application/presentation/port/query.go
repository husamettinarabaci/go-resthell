package application

import (
	"context"

	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
)

type QueryPort interface {
	IsCommandExists(ctx context.Context, commandRequest me.CommandRequest) (bool, error)
}
