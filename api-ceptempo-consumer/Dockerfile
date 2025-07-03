FROM golang:1.24.3-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o consumer

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/consumer .
EXPOSE 8081
CMD ["./consumer"]