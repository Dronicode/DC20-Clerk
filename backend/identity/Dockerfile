FROM golang:1.25 AS dev
WORKDIR /app
COPY . .
CMD ["go", "run", "./cmd/main.go"]

FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go build -o identity ./cmd/main.go

FROM alpine AS prod
WORKDIR /app
COPY --from=build /app/identity .
EXPOSE 8081
CMD ["./identity"]
