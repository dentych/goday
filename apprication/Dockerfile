FROM golang:1.10-alpine

RUN apk add --no-cache git

COPY . /go/src/github.com/dentych/goday/apprication

RUN go get ./...
RUN go build github.com/dentych/goday/apprication

FROM alpine

WORKDIR /work

COPY --from=0 /go/bin/apprication /work

EXPOSE 8080

CMD /work/apprication