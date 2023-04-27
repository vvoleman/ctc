FROM golang:1.20

WORKDIR /src

COPY go.mod main.go ./

RUN go build -o app .
RUN chmod +x app

EXPOSE 8080

CMD ["./app"]