package domain

import tjson "github.com/husamettinarabaci/go-resthell/tool/json"

type Command struct {
	Values []string `json:"values"`
}

func (o Command) ToJson() string {
	return tjson.ToJson(o)
}

func (a Command) FromJson(i string) Command {
	return tjson.FromJson[Command](i)
}

func NewCommand(values []string) Command {
	return Command{
		Values: values,
	}
}

func (o Command) IsEmpty() bool {
	return o.ToJson() == Command{}.ToJson()
}

func (o Command) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Command) IsEqual(i Command) bool {
	return o.ToJson() == i.ToJson()
}
