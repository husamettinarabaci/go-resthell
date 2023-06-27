FROM golang:1.20-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  portal

WORKDIR /go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main ./github.com/loop-soft/portal-be-country/cmd/main.go

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .
COPY --from=base /go/github.com/loop-soft/portal-be/config/ ./config/

USER portal:portal
EXPOSE 18000
EXPOSE 18080

CMD ["./main"]