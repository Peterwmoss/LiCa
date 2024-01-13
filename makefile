all: run

get:
	go get ./cmd/server

run: postgres-start
	air

t:
	templ generate

postgres-start:
	@printf "Creating postgres container, or reusing existing\n"
	@(docker run -itd --name lica-postgres -p 5433:5432 -e POSTGRES_PASSWORD=postgres postgres &> /dev/null && sleep 1) || (docker start lica-postgres && sleep 1)
	@printf "Creating database if it does not exist\n"
	@docker exec -it lica-postgres psql -U postgres -d postgres -c "CREATE DATABASE lica;" &> /dev/null || true

postgres-stop:
	docker stop lica-postgres

postgres-rm: postgres-stop
	docker rm lica-postgres

clean:
	rm -rf ./bin
	go mod tidy
