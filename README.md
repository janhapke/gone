# gone
A minimalistic HTTP Server that retuns an empty HTML page.

Written in Go, statically compiled and built as a barebones docker container without an underlying operating system.

## Run locally

    docker run --rm -it -v ${PWD}:/data --workdir /data -p 8888:80 golang go run gone.go

Then go to http://localhost:8888

## Build Docker Container

    docker build .

## Test built image locally

    docker run --rm -it -p 8888:80 <IMAGE ID>

Then go to http://localhost:8888
