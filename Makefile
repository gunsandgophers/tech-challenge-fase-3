build:
	docker compose up -d --build

start:
	docker compose up -d

stop:
	docker compose down

logs/app:
	docker compose logs -f --no-log-prefix app 
