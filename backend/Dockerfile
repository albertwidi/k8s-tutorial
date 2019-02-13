FROM golang:1.11.5
LABEL stage=intermediate
WORKDIR $GOPATH/src/github.com/albertwidi/k8s-tutorial/backend
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -v -o backend-bin main.go

FROM alpine:latest
COPY --from=0 /go/src/github.com/albertwidi/k8s-tutorial/backend/backend-bin /backend-bin
ENTRYPOINT /backend-bin