FROM golang:1.10-alpine

RUN apk add --no-cache git

COPY . /go/src/github.com/dentych/goday/james

RUN go get ./...
RUN go build github.com/dentych/goday/james

FROM alpine

WORKDIR /work

COPY --from=0 /go/bin/james /work

EXPOSE 8081

CMD /work/james