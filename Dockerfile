FROM docker.io/golang:1.24-alpine AS build
RUN apk add --no-cache alpine-sdk

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main cmd/api/main.go

FROM docker.io/alpine:3.20.1 AS prod

WORKDIR /app
COPY --from=build /app/main /app/main

EXPOSE 8080
CMD ["./main"]

