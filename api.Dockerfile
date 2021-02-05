FROM golang:1.15.7 as builder

COPY ./api/src /go/src

WORKDIR /go/src/

RUN go install && \
    go build -o webpractice

FROM golang:1.15.7
COPY --from=builder /go/src /go/src
COPY ./wait-for-it.sh /go/src/wait-for-it.sh
WORKDIR /go/src/
