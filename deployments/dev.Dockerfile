FROM golang:1.13-alpine
MAINTAINER "Igor Sant'ana <contato@igorsantana.com>"

RUN apk --no-cache --update add dep git

RUN mkdir -p /go/src

ENV GOPATH /go
ENV GOBIN /go/bin