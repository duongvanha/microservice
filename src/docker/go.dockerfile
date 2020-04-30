FROM golang:1.13.3 as builder

ARG dir
ENV GO111MODULE=on

COPY . /go/src/
RUN cd /go/src/ && go mod download
WORKDIR /go/src/${dir}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service

FROM alpine:3.7
COPY --from=builder /service /
WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["/service"]
