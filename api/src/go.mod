module github.com/cheesecat47/webpractice/api

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/cheesecat47/webpractice/constant v0.0.0
	github.com/cheesecat47/webpractice/model v0.0.0
	github.com/cheesecat47/webpractice/server v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
)

replace (
	github.com/cheesecat47/webpractice/constant v0.0.0 => ./constant
	github.com/cheesecat47/webpractice/model v0.0.0 => ./model
	github.com/cheesecat47/webpractice/server v0.0.0 => ./server
)
