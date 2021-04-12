FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

EXPOSE 5000

RUN go build

CMD ["./RestServer-Go"]

