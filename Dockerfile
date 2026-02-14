FROM golang:1.24-alpine AS builder
ENV GOARCH=amd64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o inapppaymentservice

FROM alpine:3.20
RUN apk --no-cache add ca-certificates wget
WORKDIR /app
COPY --from=builder /app/inapppaymentservice .
COPY --from=builder /app/.env .
EXPOSE 8103
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget -qO- http://localhost:8103/health || exit 1
CMD ["./inapppaymentservice"]
