# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# 生成sql代码
sqlc:
	sqlc generate

# 启动postgres容器
postgres-up:
	docker compose -f postgres-compose.yml up -d

# 停止postgres容器
postgres-down:
	docker compose -f postgres-compose.yml down

# 创建simple_bank数据库
postgres-create-db:
	docker exec -it postgres15 createdb --username postgres --owner postgres simple_bank
	#docker exec -it postgres15 psql simple_bank --username postgres

# 删除simple_bank数据库
postgres-drop-db:
	docker exec -it postgres15 dropdb simple_bank --username postgres

# 升级迁移, 先安装https://github.com/golang-migrate/migrate/tree/master
migrate-up:
	migrate -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -path db/migrate -verbose up

# 降级迁移文件, 先安装https://github.com/golang-migrate/migrate/tree/master
migrate-down:
	migrate -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -path db/migrate -verbose down

# go test
test:
	go test -v -cover ./...

postgres-first:
	make postgres-up && make postgres-create-db && make migrate-up

# restart
postgres-restart:
	make migrate-down && make migrate-up

server:
	go run main.go

.PHONY: sqlc postgres-up postgres-down postgres-create-db postgres-drop-db migrate-up migrate-down test server
