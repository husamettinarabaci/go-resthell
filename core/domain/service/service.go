package domain

import (
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (a Service) IsCommandRequestEntityValid(commandRequest me.CommandRequest) error {
	return commandRequest.IsValid()
}
