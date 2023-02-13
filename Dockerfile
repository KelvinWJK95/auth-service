FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /auth-service

EXPOSE 8080

ENTRYPOINT [ "/auth-service" ]