# golang-todo-app

docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine 

docker exec -it todopostgres12 psql -U root 

docker logs postgres12 

golang-migrate
