FROM golang:1.16.3

WORKDIR /go/src/api-test
COPY . .
EXPOSE 8080
RUN go build .
CMD ["./api-test"]