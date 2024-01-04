FROM golang:1.21.5-alpine

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8081

CMD ["./main", "start-server"]