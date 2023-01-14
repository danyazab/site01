gener: swag init -g cmd/main.go

generate-docs:
		@swag fmt -d
		@swag init -d site01/internal -o site01/docs --parseDependency