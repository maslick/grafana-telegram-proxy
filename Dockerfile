FROM golang:alpine as builder
RUN apk add --no-cache ca-certificates git

WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch as runtime
COPY --from=builder /src/grafana-telegram-proxy ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/grafana-telegram-proxy"]