package domain

import (
	"strings"

	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-resthell/tool/json"
)

type CommandRequest struct {
	Id      uuid.UUID  `json:"id"`
	Command mo.Command `json:"command"`
}

func (e CommandRequest) ToJson() string {
	return tjson.ToJson(e)
}

func (a CommandRequest) FromJson(i string) CommandRequest {
	return tjson.FromJson[CommandRequest](i)
}

func NewCommandRequest(id uuid.UUID, command mo.Command) CommandRequest {
	return CommandRequest{
		Id:      id,
		Command: command,
	}
}

func (o CommandRequest) IsEmpty() bool {
	return o.ToJson() == CommandRequest{}.ToJson()
}

func (o CommandRequest) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o CommandRequest) IsEqual(i CommandRequest) bool {
	return o.ToJson() == i.ToJson()
}

func (a CommandRequest) IsValid() error {
	if a.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if a.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if a.Command.IsEmpty() {
		return mo.ErrInvalidInput
	}
	for _, v := range a.Command.Values {
		if len(v) == 0 {
			return mo.ErrInvalidInput
		}
		if strings.Contains(v, "|") {
			return mo.ErrPipeIsNotSupported
		}
	}

	return nil
}
