build:
	docker compose build
run:
	docker compose up -d
prune:
	docker system prune --all