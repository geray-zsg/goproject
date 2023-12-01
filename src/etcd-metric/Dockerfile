FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY main.go .main.go

RUN go build -o etcd-metrics-exporter .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/etcd-metrics-exporter .

ENV ETCD_BACKUP_FILE_PATH=/tmp/etcd_backup_status.prom

EXPOSE 8080

CMD ["./etcd-metrics-exporter"]

