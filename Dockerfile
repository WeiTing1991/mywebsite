FROM golang:1.23.0-alpine
WORKDIR /mywebsite

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./build/myweb .

EXPOSE 8080

CMD ["./build/myweb"]

