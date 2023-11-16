DB_URL=postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgresql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:14-alpine

createdb:
	docker exec -it postgresql createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresql dropdb simple_bank

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test