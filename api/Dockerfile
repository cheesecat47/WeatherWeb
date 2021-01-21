FROM golang:1.15.7 as builder

COPY ./src /go/src

WORKDIR /go/src/

RUN go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql

RUN go build

FROM golang:1.15.7
COPY --from=builder /go/src /go/src
CMD ["/go/src/webpractice"]