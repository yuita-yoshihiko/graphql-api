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
