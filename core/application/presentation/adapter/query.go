package application

import (
	"context"

	as "github.com/husamettinarabaci/go-resthell/core/application/service"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
)

type QueryAdapter struct {
	service as.Service
}

func NewQueryAdapter(s as.Service) QueryAdapter {
	return QueryAdapter{
		service: s,
	}
}

func (a QueryAdapter) IsCommandExists(ctx context.Context, commandRequest me.CommandRequest) (bool, error) {
	return a.service.IsCommandExists(ctx, commandRequest)
}
