migrate-db-up:
	@go run cmd/cli/cli.go db:migrate --up
swagger-doc:
	@swag init -g server.go --output docs/swagger --parseDependency