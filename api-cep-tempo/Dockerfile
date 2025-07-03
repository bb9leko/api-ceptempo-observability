FROM golang:1.24 as build
WORKDIR /app
COPY . .
WORKDIR /app/cmd/api-cep-tempo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/api-cep-tempo 

FROM alpine
WORKDIR /app
COPY --from=build /app/api-cep-tempo .
COPY ./static ./static
ENTRYPOINT ["./api-cep-tempo"]