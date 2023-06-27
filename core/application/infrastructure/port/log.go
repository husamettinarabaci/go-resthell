package application

import (
	"context"

	mi "github.com/husamettinarabaci/go-resthell/core/domain/model/interface"
)

type LogPort interface {
	Log(ctx context.Context, source string, logData mi.Loggable)
}
