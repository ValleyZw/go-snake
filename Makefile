default: test

build:
	go build -v -o ./bin/go-snake

run: build
	./bin/go-snake

docker_build:
	docker build -t valleyzw/go-snake .

docker_run:
	docker run -it valleyzw/go-snake

clear:
	rm -rf ./bin

test:
	go test -v ./...	