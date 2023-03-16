ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /firead
WORKDIR /firead/string-id/integration
RUN go mod download
RUN go build -o ./firead-string-id ./main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder ./firead/string-id/integration/firead-string-id ./
EXPOSE 8080
ENTRYPOINT ["./firead-string-id"]