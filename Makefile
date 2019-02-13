compose-up:
	@docker-compose build
	@docker-compose up -d

compose-down:
	@docker-compose down