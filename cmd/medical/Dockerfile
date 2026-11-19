FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o medical ./cmd/medical

ENV SERVICE_ID=medical-1
ENV SERVICE_NAME=medical

EXPOSE 8082 50052

CMD ["./medical", "-config=./resources/config.yaml"]