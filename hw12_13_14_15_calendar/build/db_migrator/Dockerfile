FROM golang:1.19 as build
LABEL stage=builder

RUN go install github.com/pressly/goose/cmd/goose@latest

WORKDIR /app

RUN go build -o cmd/db_migrator/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/db_migrator /app/db_migrator

ENV CONFIG_FILE /etc/calendar/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

CMD ./db_migrator -config ${CONFIG_FILE}