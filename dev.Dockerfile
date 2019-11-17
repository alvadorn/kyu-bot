FROM golang:1.13-alpine
MAINTAINER "Igor Sant'ana <contato@igorsantana.com>"

RUN adduser -D -u 1000 -g kyu_bot kyu_bot

RUN apk --no-cache --update add git

RUN mkdir -p /home/app

WORKDIR /home/app
