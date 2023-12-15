FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/review-api
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/review-api /go/src/review-api


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/review-api /usr/local/bin/review-api
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["review-api"]
