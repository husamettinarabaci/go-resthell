package domain

import tjson "github.com/husamettinarabaci/go-resthell/tool/json"

type Response struct {
	StdOut []string `json:"stdout"`
	StdErr []string `json:"stderr"`
}

func (o Response) ToJson() string {
	return tjson.ToJson(o)
}

func (a Response) FromJson(i string) Response {
	return tjson.FromJson[Response](i)
}

func NewResponse(stdOut []string, stdErr []string) Response {
	return Response{
		StdOut: stdOut,
		StdErr: stdErr,
	}
}

func (o Response) IsEmpty() bool {
	return o.ToJson() == Response{}.ToJson()
}

func (o Response) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Response) IsEqual(i Response) bool {
	return o.ToJson() == i.ToJson()
}
