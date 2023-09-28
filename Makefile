init: \
    build \
	up \
	migrate_up

build:
	docker build -t skeleton_schema_generator docker/tools/schema_generator
up:
	docker-compose up -d
down:
	docker-compose down
migrate_up:
	sh docker/shell/recreate-db.sh