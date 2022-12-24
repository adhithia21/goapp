FROM golang:alpine3.16 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /app

FROM golang:alpine3.16
WORKDIR /
COPY --from=build /app /app
EXPOSE 80
ENTRYPOINT ["/app"]