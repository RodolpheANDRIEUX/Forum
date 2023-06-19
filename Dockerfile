FROM alpine:latest as builder
RUN apk add --no-cache git
RUN apk add --no-cache go
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest as runner
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY web /app/web
EXPOSE 3000
CMD ["./main"]
