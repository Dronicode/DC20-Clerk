FROM golang:1.25 AS dev
WORKDIR /app
COPY . .
COPY common/ ../common/
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", "air.toml"]

FROM golang:1.25 AS build
WORKDIR /app
COPY . .
COPY common/ ../common/
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o gateway ./cmd/main.go

FROM alpine AS prod
WORKDIR /app
COPY --from=build /app/gateway .
EXPOSE 8080
CMD ["./gateway"]
