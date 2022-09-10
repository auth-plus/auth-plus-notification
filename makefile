.PHONY: infra/up
infra/up:
	docker-compose up -d database prometheus grafana

.PHONY: infra/down
infra/down:
	docker-compose down

.PHONY: dev
dev:
	make infra/up
	docker-compose up -d api
	docker-compose exec api sh

.PHONY: test/ci
test/ci:
	make infra/up
	docker-compose up -d api
	docker-compose exec -T api go test ./... -coverpkg=./... -coverprofile=c.out 
	make clean/docker

.PHONY: clean/docker
clean/docker:
	make infra/down
	docker container prune -f
	docker volume prune -f
	docker image prune -f
	rm -rf db/schema.sql
	rm -f db/schema.sql