run:
	docker-compose up --build

test:
	go clean -testcache
	go test -v ./...