infra/up:
	docker-compose up -d database database-migration

infra/down:
	docker-compose down

dev:
	make infra/up
	docker-compose up -d api
	docker-compose exec api sh

test/ci:
	make infra/up
	docker-compose up -d api
	docker-compose exec -T api go test
	make clean/docker

clean/node:
	rm -rf node_modules
	rm package-lock.json

clean/docker:
	make infra/down
	docker container prune -f
	docker volume prune -f
	docker image prune -f
	rm -rf db/schema.sql
	rm -f db/schema.sql