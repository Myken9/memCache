
init:
	docker-compose up -d

.PHONY: vendor
vendor:
	go mod vendor
	go mod tidy
