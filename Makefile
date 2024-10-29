postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=enku0811 -d postgres
postgres:
	docker exec -it postgres psql -U postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres users
dropdb:
	docker exec -it postgres dropdb goc
createtable:
	migrate create -ext sql -dir db/migrate -seq event_quota_table
migrateup:
	migrate -path db/migrate -database "postgresql://postgres:enku0811@localhost:5433/users?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://postgres:enku0811@localhost:5433/users?sslmode=disable" -verbose down
# when there is dirty migration, use force to and migrateup
remove:
	migrate -path db/migrate -database "postgresql://postgres:enku0811@localhost:5433/users?sslmode=disable" force versionnumber

.PHONEY: postgresinit postgres createdb dropdb migrateup migratedown remove createtable