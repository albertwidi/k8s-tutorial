FROM golang:1.11.5

WORKDIR $GOPATH/src/github.com/albert/k8s-tutorial/backend
COPY . .

RUN go build -v -o bin-backend main.go