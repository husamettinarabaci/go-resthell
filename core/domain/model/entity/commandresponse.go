package domain

import (
	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-resthell/tool/json"
)

type CommandResponse struct {
	Id       uuid.UUID   `json:"id"`
	Response mo.Response `json:"response"`
}

func (e CommandResponse) ToJson() string {
	return tjson.ToJson(e)
}

func (a CommandResponse) FromJson(i string) CommandResponse {
	return tjson.FromJson[CommandResponse](i)
}

func NewCommandResponse(id uuid.UUID, response mo.Response) CommandResponse {
	return CommandResponse{
		Id:       id,
		Response: response,
	}
}

func (o CommandResponse) IsEmpty() bool {
	return o.ToJson() == CommandResponse{}.ToJson()
}

func (o CommandResponse) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o CommandResponse) IsEqual(i CommandResponse) bool {
	return o.ToJson() == i.ToJson()
}

func (a CommandResponse) IsValid() error {
	if a.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if a.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if a.Response.IsEmpty() {
		return mo.ErrInvalidInput
	}
	return nil
}
