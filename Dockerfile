FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build cmd/gocrud/*.go

EXPOSE 8000

CMD ["./cli_cmd","start"]