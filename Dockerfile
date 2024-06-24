#build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk add -U tzdata
RUN apk --update add ca-certificates
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o app ./cmd/app
RUN chmod 777 app

#final stage
FROM scratch
WORKDIR /res
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/app .
COPY --from=builder /app/.env .
LABEL Name=paperidtest Version=0.0.1
EXPOSE 1323

ENTRYPOINT ["./app"]