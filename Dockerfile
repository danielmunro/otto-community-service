FROM golang:1.17
WORKDIR /go/src
COPY . .
RUN go get -d -v ./...
RUN go build
EXPOSE 8081
ENTRYPOINT ["./otto-community-service"]
