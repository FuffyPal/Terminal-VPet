FROM docker.io/library/golang:1.26.3-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY src/ ./src/

RUN CGO_ENABLED=0 GOOS=linux go build -o pet ./src

FROM scratch

ENV HOME=/root

WORKDIR /root/

COPY --from=builder /app/pet .

ENTRYPOINT ["./pet"]
