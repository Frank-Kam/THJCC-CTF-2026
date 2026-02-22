FROM golang:1.21-alpine

WORKDIR /app

COPY server.go /tmp/server.go

RUN go build -o /app/gojail_server /tmp/server.go

RUN adduser -D user

RUN chmod 755 /app/gojail_server && \
    chown user:user /app/gojail_server

RUN rm /tmp/server.go

USER user

EXPOSE 9000

CMD ["/app/gojail_server"]
