postgres:
	docker run --name todopostgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it todopostgres12 createdb --username=root --owner=root todo_db  

dropdb:
	docker exec -it todopostgres12 dropdb todo_db  

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable" -verbose up  

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable" -verbose down 

.PHONY: postgres createdb dropdb migrateup migratedown