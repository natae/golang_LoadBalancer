FROM golang:1.18-alpine3.16

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

WORKDIR /app/example

RUN go build -o /test-lb

EXPOSE 3000

CMD [ "/test-lb" ]