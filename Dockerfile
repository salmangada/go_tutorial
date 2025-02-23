# Build stage
FROM golang:1.23-alpine3.19 AS builder
WORKDIR /app
COPY . .
ENV GOTOOLCHAIN=auto
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

RUN chmod +x /app/start.sh /app/wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]