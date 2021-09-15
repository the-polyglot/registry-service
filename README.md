# migrate database 
# $ migrate create -ext sql -dir db/migrations -seq create_items_table
# $ export POSTGRESQL_URL="postgres://bucketeer:bucketeer_pass@localhost:5432/registry?sslmode=disable"
# $ migrate -database ${POSTGRESQL_URL} -path db/migrations