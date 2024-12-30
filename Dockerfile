FROM golang:1.22.5-alpine3.20 as builder

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

RUN update-ca-certificates

FROM alpine:3.20

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY school /

CMD ["/school"]
