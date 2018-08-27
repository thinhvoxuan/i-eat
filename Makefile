ARGS = $(filter-out $@,$(MAKECMDGOALS))

build:
	docker-compose build

dev_up:
	docker-compose up -d

down:
	docker-compose down

dep: 
	docker-compose exec app dep $(ARGS)

dep-status: 
	make dep status

dev-local: 
	realize start

%:
	@: