ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine3.20 as builder

ARG BUILD_CMD

WORKDIR /app

COPY . .

RUN go mod download && sh -c "${BUILD_CMD}"

# RUN apk update && apk upgrade && apk add --no-cache ca-certificates

RUN update-ca-certificates

FROM alpine:3.20

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates

COPY --from=builder /app/school-bot .

CMD ["/app/school-bot"]

