# Stage 1: Dev mode (with Go toolchain)
FROM golang:1.25 AS dev
WORKDIR /app
COPY . .
CMD ["go", "run", "./cmd/main.go"]

# Stage 2: Build binary for prod
FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go build -o gateway ./cmd/main.go

# Stage 3: Prod mode (minimal runtime)
FROM alpine AS prod
WORKDIR /app
COPY --from=build /app/gateway .
EXPOSE 8080
CMD ["./gateway"]
