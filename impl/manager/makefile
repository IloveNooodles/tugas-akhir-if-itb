migrate-new:
	migrate create -seq -ext sql -dir migrations "${name}";

migrate-up:
	go run main.go migrate-up

migrate-down:
	go run main.go migrate-down

manager:
	go run main.go start

docker:
	docker compose up -d
