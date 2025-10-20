# Stage 1: Build
FROM golang:1.25 AS builder
WORKDIR /src
COPY . .
RUN go mod tidy && go build -o /app .

# Stage 2: Run
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
