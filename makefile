dev:
	export POSTGRESQL_URL="postgres://registry:registry_pass@localhost:5432/registry?sslmode=disable"
	migrate -database ${POSTGRESQL_URL} -path db/migrations up