FROM golang as builder

COPY ./src /go/src

WORKDIR /go/src/

RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql

FROM golang
COPY --from=builder /go/src /go/src
CMD ["go", "run", "/go/src/main.go"]