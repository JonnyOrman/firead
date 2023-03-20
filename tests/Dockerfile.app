ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS builder

ARG ID_TYPE
ARG APP_TYPE

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /firead
WORKDIR /firead/${ID_TYPE}-id/integration-${APP_TYPE}
RUN go mod download
RUN go build -o ./app ./main.go

FROM alpine:latest

ARG ID_TYPE
ARG APP_TYPE

WORKDIR /root/
COPY --from=builder ./firead/${ID_TYPE}-id/integration-${APP_TYPE}/app ./
EXPOSE 8080
ENTRYPOINT ["./app"]