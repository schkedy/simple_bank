postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=2356 -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:2356@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:2356@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:2356@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:2356@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go simple_bank/db/sqlc Store
test:
	go test -v -cover -coverprofile=coverage.out ./...

server:
	go run main.go

covershow:
	go tool cover -html=coverage.out -o coverage.html && firefox coverage.html

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test covershow mock migrateup1 migratedown1