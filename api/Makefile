migrate:
	migrate -source file://migrations \
			-database postgres://postgres:root@127.0.0.1:5432/postgres?sslmode=disable up

migrate-down:
	migrate -source file://migrations \
			-database postgres://postgres:root@127.0.0.1:5432/postgres?sslmode=disable down