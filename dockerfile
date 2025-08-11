FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY . ./

RUN go build -o sandbox ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/sandbox .

EXPOSE 8080

CMD ["./sandbox"]