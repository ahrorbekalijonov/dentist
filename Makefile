run:
	go run cmd/main.go
	
swag:
	swag init -g api/router.go -o api/docs

migrate-up:
	migrate -path migrations -database "postgresql://postgres:0@localhost:5432/doctordb?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:0@localhost:5432/doctordb?sslmode=disable" -verbose down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq doctor

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:0@localhost:5432/doctordb?sslmode=disable" force 
