.PHONY : clearscr fresh clean all
# variables
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))

restart:
	docker restart gosite

dev-env: build run

build:
	docker build -t gositebase .

run:
	docker run -dith gosite --name gosite -p 8088:8088 -v ${mkfile_dir}:/srv gositebase
	sleep 5
	open http://0.0.0.0:8088

destroy:
	docker rm -f gosite
	docker rmi gositebase

connect:
	docker exec -it gosite ash

log:
	docker logs -f gosite
