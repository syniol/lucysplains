up:
	docker-compose -f ./deploy/docker-compose.yml up -d

down:
	docker-compose -f ./deploy/docker-compose.yml down

tty:
	docker-compose -f ./deploy/docker-compose.yml exec golang sh
