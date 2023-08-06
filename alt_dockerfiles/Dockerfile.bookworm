FROM golang:1.20.7

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-web-app

EXPOSE 8080

CMD ["/hello-web-app"]
