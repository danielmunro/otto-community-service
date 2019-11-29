FROM golang:1.13
WORKDIR /go/src/github.com/danielmunro/otto-community-service
COPY . .
RUN go get -d -v ./internal/...
RUN go build
EXPOSE 8081
ENTRYPOINT ["./otto-community-service"]
