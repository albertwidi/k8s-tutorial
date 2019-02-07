FROM golang:1.11.5
LABEL stage=intermediate
WORKDIR $GOPATH/src/github.com/albertwidi/k8s-tutorial/backend
COPY . .
RUN GOOS=linux go build -v -o backend-bin main.go

FROM alpine:latest
WORKDIR /
COPY --from=0 /go/src/github.com/albertwidi/k8s-tutorial/backend/backend-bin .
COPY entrypoint.sh .
ENTRYPOINT /entrypoint.sh