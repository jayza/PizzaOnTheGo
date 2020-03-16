FROM golang:1.13-alpine

ENV RECEIPT_FILE_DIRECTORY=/order-receipts/

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN mkdir -p $RECEIPT_FILE_DIRECTORY

EXPOSE 8080

CMD ["./main"]