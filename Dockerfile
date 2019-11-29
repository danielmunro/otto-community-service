FROM golang:1.13.4
WORKDIR /go/src
COPY . .
RUN go get -d -v ./internal/...
RUN go build
EXPOSE 8081
ENTRYPOINT ["./otto-community-service"]
