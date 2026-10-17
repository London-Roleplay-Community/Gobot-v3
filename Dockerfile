# binaries
FROM golang:1.23-alpine AS builder

RUN apk --no-cache add git ca-certificates
WORKDIR /docker

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gobot ./src/app.go

# image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /docker/gobot /gobot

USER 65532:65532

ENTRYPOINT ["/gobot"]