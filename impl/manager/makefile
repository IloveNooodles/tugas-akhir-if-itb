migrate-new:
	# supply with name=x
	migrate create -ext sql -dir migrations "${name}";

migrate:
	go run main.go migrate-up
	