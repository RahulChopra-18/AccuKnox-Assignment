FROM golang:latest

LABEL maintainer=""

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./go-crud"]