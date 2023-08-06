FROM golang:1.20.7 AS build-stage

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /hello-web-app

FROM scratch AS build-release-stage

WORKDIR /

COPY --from=build-stage /hello-web-app /hello-web-app

EXPOSE 8080

ENTRYPOINT ["/hello-web-app"]
