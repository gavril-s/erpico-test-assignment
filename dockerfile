FROM golang:1.22.2-alpine as builder
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

RUN echo '#!/bin/sh' > ./entrypoint.sh \
    && echo 'sleep 10' >> ./entrypoint.sh \
    && echo './app' >> ./entrypoint.sh \
    && chmod +x ./entrypoint.sh

EXPOSE 5000
ENTRYPOINT ["./entrypoint.sh"]
