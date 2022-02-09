FROM golang:1.17.0-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -ldflags="-w -s" -o main .

CMD ["/app/main"]