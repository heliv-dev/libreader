FROM golang:latest

WORKDIR /work

COPY . .

RUN go mod download

RUN go build -o library-app .

EXPOSE 8080
CMD ["./library-app"]