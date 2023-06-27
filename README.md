# go-resthell
Micro-DDD RestAPI for Shell Executor

DDD
Test
Document
CI/CD
Release
Docker

Gin
Logrus
Ioc
yaml
UUID

### Project Structure
```bash
.
├── cmd
│   └── main.go
├── config
│   ├── log_local.yml
│   ├── log.yml
│   ├── rest_local.yml
│   └── rest.yml
├── core
│   ├── application
│   │   ├── infrastructure
│   │   │   └── port
│   │   │       ├── log.go
│   │   │       └── shell.go
│   │   ├── presentation
│   │   │   ├── adapter
│   │   │   │   ├── command.go
│   │   │   │   └── query.go
│   │   │   └── port
│   │   │       ├── command.go
│   │   │       └── query.go
│   │   └── service
│   │       └── service.go
│   └── domain
│       ├── model
│       │   ├── entity
│       │   │   ├── commandrequest.go
│       │   │   └── commandresponse.go
│       │   ├── interface
│       │   │   └── loggable.go
│       │   └── object
│       │       ├── command.go
│       │       ├── error.go
│       │       └── response.go
│       └── service
│           └── service.go
├── go.mod
├── go.sum
├── LICENSE
├── pkg
│   ├── infrastructure
│   │   ├── adapter
│   │   │   ├── log.go
│   │   │   └── shell.go
│   │   └── mapper
│   │       └── command.go
│   └── presentation
│       ├── controller
│       │   └── rest
│       │       └── restapi.go
│       └── dto
│           ├── commandrequest.go
│           └── commandresponse.go
├── README.md
├── script
│   ├── Dockerfile
│   └── k8s.yml
├── TODO.md
└── tool
    ├── config
    │   ├── log.go
    │   └── rest.go
    └── json
        └── json.go
```

Usage
```
export LOCAL=true && go run cmd/main.go

POST localhost:18080/api/cmd HTTP/1.1
{
    "command": ""
}
```

Docker
```
docker build -t devhsmtek/resthell -f script/Dockerfile .
docker tag devhsmtek/resthell devhsmtek/resthell:latest
docker push devhsmtek/resthell:latest
docker run -p 18080:18080 devhsmtek/resthell:latest
```

Kubernetes
```
kubectl apply -f script/k8s.yml
```

CI/CD
```
DOCKERHUB_USERNAME
DOCKERHUB_TOKEN

pull_request to release branch
```

