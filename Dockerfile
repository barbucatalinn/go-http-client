FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates tzdata git && update-ca-certificates
RUN adduser -D -g '' unprivileged_user
WORKDIR /code
COPY . .
FROM scratch as production
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER unprivileged_user
HEALTHCHECK NONE