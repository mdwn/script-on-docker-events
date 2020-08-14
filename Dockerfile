FROM golang:1.15 AS builder

WORKDIR /go/src/github.com/meowfaceman/script-on-docker-hosts

COPY . .

RUN make

FROM alpine

COPY --from=builder /go/src/github.com/meowfaceman/script-on-docker-hosts/script-on-docker-events /app/script-on-docker-events

RUN apk --no-cache add docker-cli bash libc6-compat

ENV PATH "/app:$PATH"

ENTRYPOINT ["script-on-docker-events"]
