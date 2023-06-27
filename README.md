# go-resthell
Micro-DDD RestAPI for Shell Executor




```
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

Project Structure
TODO
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

