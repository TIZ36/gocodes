# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app
# Usually the very first thing you do once
# you’ve downloaded a project written in Go is to install the modules
# necessary to compile it.
COPY go.mod go.sum ./

RUN go mod download && go mod verify

ADD . .
RUN go build -o ./gomock

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build gomock ./gomock

EXPOSE 8888

USER nonroot:nonroot

ENTRYPOINT ["./docker-gs-ping"]
