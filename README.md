# go-resthell
Micro-DDD (Hex-Arc) RestAPI for Shell Executor 

[![Build](https://github.com/husamettinarabaci/go-resthell/actions/workflows/build.yml/badge.svg)](https://github.com/husamettinarabaci/go-resthell/actions/workflows/build.yml)
[![Test](https://github.com/husamettinarabaci/go-resthell/actions/workflows/test.yml/badge.svg)](https://github.com/husamettinarabaci/go-resthell/actions/workflows/test.yml)
[![Test](https://github.com/husamettinarabaci/go-resthell/actions/workflows/sast.yml/badge.svg)](https://github.com/husamettinarabaci/go-resthell/actions/workflows/sast.yml)

<a href="https://kaos.sh/g/go-badge"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
<a href="https://kaos.sh/w/go-badge/ci"><img src="https://kaos.sh/w/go-badge/ci.svg" alt="GitHub Actions CI Status" /></a>
<a href="https://github.com/husamettinarabaci/go-resthell/stargazers"><img src="https://img.shields.io/github/stars/husamettinarabaci/go-resthell" alt="Stars Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-resthell/network/members"><img src="https://img.shields.io/github/forks/husamettinarabaci/go-resthell" alt="Forks Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-resthell/pulls"><img src="https://img.shields.io/github/issues-pr/husamettinarabaci/go-resthell" alt="Pull Requests Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-resthell/issues"><img src="https://img.shields.io/github/issues/husamettinarabaci/go-resthell" alt="Issues Badge"/></a>
<a href="https://github.com/husamettinarabaci/go-resthell/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/husamettinarabaci/go-resthell?color=2b9348"></a>
<a href="https://github.com/husamettinarabaci/go-resthell/blob/master/LICENSE"><img src="https://img.shields.io/github/license/husamettinarabaci/go-resthell?color=2b9348" alt="License Badge"/></a>

## Stack
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)

## Dependencies
[![Go-Gin](https://img.shields.io/badge/GoLib-Gin-green.svg)](https://github.com/gin-gonic/gin/)
[![Go-Ioc](https://img.shields.io/badge/GoLib-Ioc-green.svg)](https://github.com/golobby/container/v3/)
[![Go-Json](https://img.shields.io/badge/GoLib-Json-green.svg)](https://github.com/goccy/go-json/)
[![Go-Yaml](https://img.shields.io/badge/GoLib-Yaml-green.svg)](https://gopkg.in/yaml.v3/)
[![Go-Uuid](https://img.shields.io/badge/GoLib-Uuid-green.svg)](https://github.com/google/uuid/)
[![Testify](https://img.shields.io/badge/GoLib-Testify-green.svg)](https://github.com/stretchr/testify/)

## Getting Started
<b>Resthell</b> provides the ability to run a shell command with RestAPI. You can run it locally or as a container. 

It has been developed with <b>Domain Driven Design (Hex-Arc)</b> architecture and allows you to be included in the domain and perform external operations in all microservice infrastructures without additional development processes. 

Do you want to learn more information about <b>Domain Driven Design</b> and <b>Hex-Arc</b>? 
 - [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
 - [Hex-Arc](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

Also, you can see these pictures about <b>Domain Driven Design (Hex-Arc)</b>. 
 - [Hex-Arc-1](https://github.com/husamettinarabaci/go-resthell/tree/main/doc/Hex-Arc-1.jpg)
 - [Hex-Arc-2](https://github.com/husamettinarabaci/go-resthell/tree/main/doc/Hex-Arc-2.jpg)


## Installation
```bash
git clone https://github.com/husamettinarabaci/go-resthell.git
cd go-resthell
go mod download
```

## Test
```bash
go test -v ./...
```

## Usage
```bash
POST https://localhost:18080/api/cmd HTTP/1.1
content-type: application/json

{
    "command": "ls -al"
}
```

## Local Run
```bash
export LOCAL=true && go run cmd/main.go
```

## Docker Build & Run
```bash
docker build -t {DOCKER_USERNAME}/{YOUR_REPO} -f script/Dockerfile .
docker tag {DOCKER_USERNAME}/{YOUR_REPO} {DOCKER_USERNAME}/{YOUR_REPO}:latest
docker push {DOCKER_USERNAME}/{YOUR_REPO}:latest
docker run -p 18080:18080 {DOCKER_USERNAME}/{YOUR_REPO}:latest
```

## Kubernetes Deploy
```bash
kubectl apply -f script/k8s.yml
```

## Github Actions
Fork the project and create the below secrets in your repo.

```bash
DOCKERHUB_USERNAME

DOCKERHUB_TOKEN
```

Create "release" branch and create a pull request to "release" branch and merge it. Github Actions will build and push docker image to your dockerhub repo.

## Project Structure - Domain Driven Design (Hex-Arc)
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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Tasks
- [ ] Implement Pipe Supporting

## Contact

Hüsamettin ARABACI - info@husamettinarabaci.com

Project Link: [https://github.com/husamettinarabaci/go-resthell](https://github.com/husamettinarabaci/go-resthell)

