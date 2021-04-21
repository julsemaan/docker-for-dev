FROM golang:1.14.3-alpine AS build
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /out/supertest .

FROM alpine AS bin

RUN apk update
RUN apk add postgresql-client

COPY --from=build /out/supertest /


