FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 80

RUN go build

CMD ["./golang-web-app"]