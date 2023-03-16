ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /firead
WORKDIR /firead/int-id/integration
RUN go mod download
RUN go build -o ./firead-int-id ./main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder ./firead/int-id/integration/firead-int-id ./
EXPOSE 8080
ENTRYPOINT ["./firead-int-id"]