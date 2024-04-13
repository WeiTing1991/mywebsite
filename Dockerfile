FROM golang:1.22.0-alpine
WORKDIR /mywebsite

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./build/myweb .
CMD ./build/myweb
