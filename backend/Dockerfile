FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o scheduled-job .

CMD ["./scheduled-job"]
