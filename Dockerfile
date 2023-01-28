FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine:latest
LABEL traefik.http.routers.my-container.rule=Host(`example.com`)
WORKDIR /root/
COPY --from=builder ./app/app .
EXPOSE 8080
CMD ["./app"]