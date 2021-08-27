.PHONY : clearscr fresh clean all
# variables
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))

restart:
	docker restart gosite

dev-env: build run

build:
	docker build -t gositebase .

build-arm:
	docker build -t gositebase-arm . -f homelab-Dockerfile

run:
	docker run -dith gosite --name gosite -p 8088:8080 -v ${mkfile_dir}:/srv gositebase
	sleep 5
	open http://0.0.0.0:8088

run-arm:
	docker run -dith gosite-arm --name gosite-arm -p 8089:8080 -v ${mkfile_dir}:/srv gositebase-arm
	sleep 5
	#open http://0.0.0.0:8089

destroy:
	docker rm -f gosite
	docker rmi gositebase

destroy-arm:
	docker rm -f gosite-arm
	docker rmi gositebase-arm

connect:
	docker exec -it gosite ash

log:
	docker logs -f gosite
