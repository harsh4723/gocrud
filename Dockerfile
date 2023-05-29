FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build cmd/gocrud/main.go

EXPOSE 8000

CMD ["./main"]