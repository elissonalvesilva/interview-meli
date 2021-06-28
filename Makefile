.PHONY: build start stop logs-api logs-analyzer logs-tail tests-docker tests-local

CONTAINER_NAME_API := api
CONTAINER_NAME_ANALYZER := analyzer

build:
	docker-compose build

start:
	docker-compose up -d

stop:
	docker-compose down

logs-api:
	docker logs -f $(CONTAINER_NAME_API)

logs-analyzer:
	docker logs -f $(CONTAINER_NAME_ANALYZER)

logs-tail:
	docker logs -f --tail 100 $(CONTAINER_NAME_API) \
	docker logs -f --tail 100 $(CONTAINER_NAME_ANALYZER)

tests-local:
	go test --coverprofile=coverage.out ./...

show-coverage-html:
	go tool cover -html=coverage.out