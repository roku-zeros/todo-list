FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY lib lib
COPY services/task ./services/task

RUN go build -ldflags="-s -w" -o ./bin/task ./services/task/cmd

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/bin/task ./bin/task

EXPOSE 8080

CMD ["./bin/task"]
