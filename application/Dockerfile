FROM golang:1.22 as base

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main .

# Deploy stage
FROM gcr.io/distroless/base

WORKDIR /

COPY --from=base /app/main .

EXPOSE 8080

CMD ["./main"]