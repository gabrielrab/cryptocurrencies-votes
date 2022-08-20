FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/cryptocurrencies-votes

EXPOSE 3000

CMD [ "bin/cryptocurrencies-votes" ]