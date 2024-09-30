export DB_URL=postgres://tech_challenge_fase_3:JV4Nv1iBaw9uns4@tech-challenge-fase-3-rds.cdmsim2u8a3c.us-east-1.rds.amazonaws.com:5432/tech_challenge_fase_3

build:
	docker compose up -d --build

build/login:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 249752625560.dkr.ecr.us-east-1.amazonaws.com

build/prod:
	docker buildx build -t tech-challenge-go-app .

build/tag:
	docker tag gunsandgophers/tech-challenge-go-app:latest 249752625560.dkr.ecr.us-east-1.amazonaws.com/gunsandgophers/tech-challenge-go-app:latest

build/push:
	docker push 249752625560.dkr.ecr.us-east-1.amazonaws.com/gunsandgophers/tech-challenge-go-app:latest

start:
	docker compose up -d

stop:
	docker compose down

logs/app:
	docker compose logs -f --no-log-prefix app

migrate:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=./migrations/ -database ${DB_URL} up

migrate/create:
	docker run -v ./migrations:/migrations --network host migrate/migrate create -ext sql -dir ./migrations $(name)

swagger:
	docker run --rm -v ./:/code ghcr.io/swaggo/swag:latest init

swagger-mac:
	docker run --platform linux/amd64 --rm -v ./:/code ghcr.io/swaggo/swag:latest init
