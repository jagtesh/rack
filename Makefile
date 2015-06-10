all: build

build:
	docker build -t convox/agent .

test: build
	docker run -v /var/lib/boot2docker:/host/boot2docker --env-file .env convox/agent -log /host/boot2docker/docker.log -cwgroup test -cwstream test -tick 2 testapp testps i-0000000

vendor:
	godep save -r -copy=true ./...
