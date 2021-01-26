FROM golang:1.15.7 as builder

COPY ./src /go/src

WORKDIR /go/src/

RUN go get -u github.com/gin-gonic/gin && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/mysql

RUN go build

FROM golang:1.15.7
COPY --from=builder /go/src /go/src
CMD ["/go/src/webpractice"]