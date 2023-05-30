postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres createdb --username=postgres --owner=postgres links

dropdb:
	sudo docker exec -it  postgres dropdb --username=postgres links

migrateup:
	migrate -path internal/repository/migration -database "postgresql://postgres:postgres@localhost:5432/links?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/repository/migration -database "postgresql://postgres:postgres@localhost:5432/links?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown