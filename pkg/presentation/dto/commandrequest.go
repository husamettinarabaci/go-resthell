package infrastructure

import (
	"strings"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-resthell/tool/json"
)

type CommandRequest struct {
	Command string `json:"command"`
}

func (a CommandRequest) ToJson() string {
	return tjson.ToJson(a)
}

func (e CommandRequest) FromJson(i string) CommandRequest {
	return tjson.FromJson[CommandRequest](i)
}

func NewCommandRequest(command string) CommandRequest {
	return CommandRequest{
		Command: command,
	}
}

func (o CommandRequest) IsEmpty() bool {
	return o.ToJson() == CommandRequest{}.ToJson()
}

func (a CommandRequest) ToCommandRequestEntity() me.CommandRequest {
	return me.NewCommandRequest(
		uuid.New(),
		mo.NewCommand(
			strings.Split(a.Command, " "),
		),
	)
}
