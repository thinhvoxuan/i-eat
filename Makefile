ARGS = $(filter-out $@,$(MAKECMDGOALS))

dev_up:
	docker-compose up -d

down:
	docker-compose down

dep: 
	docker-compose run app dep $(ARGS)

dev-local: 
	realize start