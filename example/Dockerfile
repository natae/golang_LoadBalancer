FROM golang:1.18-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/example

RUN go build -o /test-lb

EXPOSE 3000

CMD [ "/test-lb" ]