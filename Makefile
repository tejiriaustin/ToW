api:
	go run main.go api

listener:
	go run main.go listener

push:
	golangci-lint run && git push origin main

lint:
	golangci-lint run

service:
	docker-compose up --build tow_service