FROM golang:1.24.2-alpine AS base
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build src/main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=base /app/main .
EXPOSE 8080
CMD ["./main"]