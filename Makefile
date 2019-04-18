FRONTEND_APPNAME=frontend-bin
BACKEND_APPNAME=backend-bin

compose-up:
	@docker-compose build
	@docker-compose up -d

compose-down:
	@docker-compose down

## this is for building image manually

build-image:
	make build-image-backend
	make build-image-frontend

push-image:
	@docker push albertwidi/k8s-tutorial-backend:$(tag_backend)
	@docker push albertwidi/k8s-tutorial-frontend:$(tag_frontend)

build-image-backend:
	@cd backend && GOOS=linux CGO_ENABLED=0 go build -v -o $(BACKEND_APPNAME) *.go
	@cd backend && docker build -f Dockerfile -t albertwidi/k8s-tutorial-backend . --rm

build-image-frontend:
	@ cd frontend && GOOS=linux CGO_ENABLED=0 go build -o $(FRONTEND_APPNAME) -v main.go
	@docker build -f ./frontend/Dockerfile -t albertwidi/k8s-tutorial-frontend . --rm