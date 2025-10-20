FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /out/myservice .

FROM alpine:3.18
COPY --from=builder /out/myservice /usr/local/bin/myservice
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/myservice"]
