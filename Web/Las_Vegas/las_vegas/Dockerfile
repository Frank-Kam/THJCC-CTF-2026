FROM golang:1.25.3-alpine AS build

WORKDIR /app

COPY main.go .

RUN go build -o las-vegas main.go

FROM alpine:latest
WORKDIR /app

COPY --from=build /app/las-vegas .

EXPOSE 14514

CMD ["./las-vegas"]
