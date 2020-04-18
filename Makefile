start:
	docker-compose up -d db
	docker-compose up markr

test:
	docker-compose up tests