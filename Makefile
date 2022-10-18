DB_URL=postgresql://root:secret@localhost:5432/tgram_subs?sslmode=disable
TEST_DB_URL=postgresql://root:secret@localhost:5432/tgram_subs_test?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root tgram_subs && docker exec -it postgres createdb --username=root --owner=root tgram_subs_test

dropdb:
	docker exec -it postgres dropdb tgram_subs && docker exec -it postgres dropdb tgram_subs_test

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up && migrate -path db/migration -database "$(TEST_DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down && migrate -path db/migration -database "$(TEST_DB_URL)" -verbose down
boil:
	sqlboiler psql

.PHONY: postgres createdb dropdb migrateup migratedown boil sqlc
