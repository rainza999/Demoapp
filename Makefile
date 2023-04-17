docker_up:
	docker compose up --build --force-recreate --no-deps -d api
	docker-compose up --build --force-recreate --no-deps -d mysql
	docker-compose up --build --force-recreate --no-deps -d redis
	docker-compose up --build --force-recreate --no-deps -d nginx
docker_down:
	docker-compose down --volumes
	docker container prune -f
	docker volume prune -f
	docker image prune -f
	docker network prune -f
docker_ps:
	docker container ls -a
mod_tidy:
	GO111MODULE=on go mod tidy
patch_vendor:
	GO111MODULE=on go mod vendor
unit_test:
	go test -mod=readonly ./...
