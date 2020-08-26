FROM golang:1.14.7-alpine3.12 AS builder

COPY . /src
RUN apk add build-base
RUN cd /src && go get -d && go build && cd plugin && go test -v

FROM alpine:3.12
EXPOSE 3000

ENV DRONE_DEBUG=false
ENV DRONE_ADDRESS=:3000

COPY --from=builder /src/drone-public-blocker /bin/drone-public-blocker

ENTRYPOINT ["/bin/drone-public-blocker"]
