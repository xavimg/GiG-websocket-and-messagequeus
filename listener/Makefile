run: 
	@docker rm --force /nats-main
	@docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
	go run cmd/*.go
