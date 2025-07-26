# /Dockerfile
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./...

FROM gcr.io/distroless/base
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
