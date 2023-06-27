package infrastructure

import (
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-resthell/tool/json"
)

type CommandResponse struct {
	Out []string `json:"out"`
	Err []string `json:"err"`
}

func (a CommandResponse) ToJson() string {
	return tjson.ToJson(a)
}

func (e CommandResponse) FromJson(i string) CommandResponse {
	return tjson.FromJson[CommandResponse](i)
}

func NewCommandResponse(out []string, err []string) CommandResponse {
	return CommandResponse{
		Out: out,
		Err: err,
	}
}

func (o CommandResponse) IsEmpty() bool {
	return o.ToJson() == CommandResponse{}.ToJson()
}

func FromResponseObject(response mo.Response) CommandResponse {
	return NewCommandResponse(
		response.StdOut,
		response.StdErr,
	)
}
