run:
	docker-compose build --force-rm --no-cache && docker-compose up --detach && docker-compose logs -f

