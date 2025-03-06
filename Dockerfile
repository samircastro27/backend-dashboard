FROM --platform= golang:1.23-alpine3.20 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV go env -w GO111MODULE=on

RUN apk update --no-cache && apk add --no-cache wget && apk add git && apk add bash

RUN wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN dapr init --slim
RUN dapr --version

WORKDIR /platform

COPY . .
COPY cmd/migrations cmd/migrations

RUN go mod tidy -compat=1.23 && go mod download

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]