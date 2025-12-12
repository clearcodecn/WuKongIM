build:
	docker build -t wukongim .

reload:
	@docker compose stop && docker compose rm -f && docker compose up -d
