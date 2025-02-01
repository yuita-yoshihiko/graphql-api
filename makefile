RUN=docker-compose run --service-ports --rm --workdir="/graphql-api" graphql-api

gql:
	${RUN} gqlgen

sqlboiler:
	${RUN} sh -c "sqlboiler -o domain/models --no-tests psql"

migrate/new:
	${RUN} sh -c "sql-migrate new ${FILE_NAME}"

migration/status:
	${RUN} sh -c "sql-migrate status --env='development'"

migration/up:
	${RUN} sh -c "sql-migrate up --env='development'"

migration/down:
	${RUN} sh -c "sql-migrate down --env='development'"

psql:
	psql -h 127.0.0.1 -p 5632 -U user graphql-api-db

test-db-up:
	docker compose -f ./docker-compose.test-db.yml up -d
	sleep 5
	${RUN} sh -c "sql-migrate up --env='test'"

test-db-down:
	docker compose -f ./docker-compose.test-db.yml down

output-mg:
	${RUN} sh -c "go run internal/skeleton/migration/main.go"

output-sc:
	${RUN} sh -c "go run internal/skeleton/schema/main.go"

output-uc:
	${RUN} sh -c "go run internal/skeleton/usecase/main.go"

output-cv:
	${RUN} sh -c "go run internal/skeleton/converter/main.go"

output-rp:
	${RUN} sh -c "go run internal/skeleton/repository/main.go"

output-db:
	${RUN} sh -c "go run internal/skeleton/database/main.go"

output-all-skeleton:
	make output-mg
	make output-sc
	make output-uc
	make output-cv
	make output-rp
	make output-db
