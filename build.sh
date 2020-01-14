#!/bin/sh -x

echo "build project"
CGO_ENABLED=0 go build -a -installsuffix cgo -o portScan main.go && \
echo "build project to docker image"
docker build -t port-scan . && \
echo "package docker image as a file"
docker save port-scan:latest | gzip > portScan-image.tar.gz && \
echo "packge project"
tar czvf  portScan.tar.gz portScan-image.tar.gz config_sample.json docker-compose.yml

echo "clean"
rm portScan portScan-image.tar.gz

