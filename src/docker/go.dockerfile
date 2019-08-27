FROM golang:1.10.3 as builder

ENV DIR /simple
COPY . /go/src/service
WORKDIR /go/src/service

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o /service

FROM alpine:3.7
COPY --from=builder /service /
WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["/service"]
