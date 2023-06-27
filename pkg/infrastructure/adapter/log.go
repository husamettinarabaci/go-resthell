package infrastructure

import (
	"context"
	"fmt"

	mi "github.com/husamettinarabaci/go-resthell/core/domain/model/interface"
	tconfig "github.com/husamettinarabaci/go-resthell/tool/config"
)

type LogAdapter struct {
}

func NewLogAdapter() LogAdapter {
	adapter := LogAdapter{}
	return adapter
}

func (a LogAdapter) Log(ctx context.Context, source string, logData mi.Loggable) {
	if tconfig.GetLogConfigInstance().Logger.Console {
		fmt.Println(source, " : ", logData.ToJson())
	}
}
