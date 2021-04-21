supertest:
	go build -o supertest

dev-up:
	docker-compose -f docker-compose.dev.yml up --build
