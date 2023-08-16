FROM FROM golang:1.19 as build

ENV BIN_FILE /opt/calendar/calendar-sender
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar_sender/*

FROM alpine:3.9

LABEL SERVICE="calendar_sender"