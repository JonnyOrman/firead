ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /firead
WORKDIR /firead/string-id
RUN go mod download
EXPOSE 8080
CMD ["go", "test"]