package infrastructure

import (
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-resthell/tool/json"
)

type Command struct {
	Values []string `json:"values"`
}

func (a Command) ToJson() string {
	return tjson.ToJson(a)
}

func (e Command) FromJson(i string) Command {
	return tjson.FromJson[Command](i)
}

func NewCommand(values []string) Command {
	return Command{
		Values: values,
	}
}

func FromCommandObject(command mo.Command) Command {
	return NewCommand(
		command.Values,
	)
}
