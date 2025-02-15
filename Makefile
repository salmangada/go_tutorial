postgres:
	docker run --name postgres-db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=sal123 -v pg_data:/var/lib/postgresql/data -p 5432:5432 -d postgres:17-alpine

createdb:
	docker exec -it postgres-db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-db dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:sal123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:sal123@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc