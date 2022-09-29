FROM golang:1.17.1-alpine3.14 as builder

ENV GOPROXY http://goproxy.cn

WORKDIR /app

COPY . /app/

RUN GOOS=linux go build -o main .

FROM alpine:3.13

RUN apk add ca-certificates

WORKDIR /app

COPY --from=builder /app/main /app/

COPY ./static/* /app/static

CMD ["/app/main"]
