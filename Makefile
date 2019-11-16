start:
	docker-compose up -d
	docker-compose ps
	# db
	# access -> docker exec -it postgres psql -d app_db -U app_user
	# quit -> \q # db_list -> \l # table_list -> \d
	# app
	# access -> docker exec -it go_echo_app /bin/sh

end:
	# docker-compose down --rmi local
	docker-compose down

postgres-start:
	docker-compose -f ./docker/postgres/docker-compose.postgres.yml up -d
	# db
	# access -> docker exec -it postgres-one psql -d app_db -U app_user
	# quit -> \q # db_list -> \l # table_list -> \d

postgres-end:
	docker-compose -f ./docker/postgres/docker-compose.postgres.yml down --rmi local

######======= TEST =======########

test-all:
	go test ./go ./go/infra

