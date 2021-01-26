module model

go 1.15

require (
	github.com/cheesecat47/webpractice/constant v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)

replace github.com/cheesecat47/webpractice/constant v0.0.0 => ../constant
