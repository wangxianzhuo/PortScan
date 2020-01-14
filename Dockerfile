FROM alpine

MAINTAINER xianzhuo

RUN mkdir /usr/app
WORKDIR /usr/app

COPY portScan /usr/app/portScan
COPY ./config_sample.json /usr/app/config.json

ENTRYPOINT [ "./portScan" ]