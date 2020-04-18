start:
	docker-compose up -d db
	docker-compose up markr

test:
	docker-compose up tests

dev:
	docker-compose run --service-ports --rm --entrypoint sh markr

kill:
	docker-compose down