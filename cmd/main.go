package main

import (
	"github.com/golobby/container/v3"

	pa "github.com/husamettinarabaci/go-resthell/core/application/presentation/adapter"
	as "github.com/husamettinarabaci/go-resthell/core/application/service"
	ds "github.com/husamettinarabaci/go-resthell/core/domain/service"
	ia "github.com/husamettinarabaci/go-resthell/pkg/infrastructure/adapter"
	cr "github.com/husamettinarabaci/go-resthell/pkg/presentation/controller/rest"
	tconfig "github.com/husamettinarabaci/go-resthell/tool/config"
)

var restConfig tconfig.RestConfig

func main() {
	restConfig.ReadConfig()
	var err error
	cont := container.New()

	//Domain resthell Service
	err = cont.Singleton(func() ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Resthell Shell Adapter
	err = cont.Singleton(func() ia.ShellAdapter {
		return ia.NewShellAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Resthell Log Adapter
	err = cont.Singleton(func() ia.LogAdapter {
		return ia.NewLogAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Application resthell Service
	err = cont.Singleton(func(s ds.Service, i ia.ShellAdapter, l ia.LogAdapter) as.Service {
		return as.NewService(s, i, l)
	})
	if err != nil {
		panic(err)
	}

	//Application resthell Query Adapter
	err = cont.Singleton(func(s as.Service) pa.QueryAdapter {
		return pa.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application resthell Command Adapter
	err = cont.Singleton(func(s as.Service) pa.CommandAdapter {
		return pa.NewCommandAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var commandHandler pa.CommandAdapter
	err = cont.Resolve(&commandHandler)
	if err != nil {
		panic(err)
	}

	cr.NewRestAPI(queryHandler, commandHandler).Serve(restConfig.Debug, restConfig.Restapi.Port)

}
