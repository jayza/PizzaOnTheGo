FROM golang:1.13-alpine

ENV RECEIPT_FILE_DIRECTORY=/order-receipts/
ENV SERVER_PORT=8080

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN mkdir -p $RECEIPT_FILE_DIRECTORY

EXPOSE $SERVER_PORT

CMD ["./main"]