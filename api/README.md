# webpractice - api

## Environment

| Name | Version |              Remarks               |
| :--: | :-----: | :--------------------------------: |
|  Go  | 1.15.7  |           Docker version           |
| Gin  |  1.6.3  | <https://github.com/gin-gonic/gin> |
| Gorm | 1.20.11 |      <https://gorm.io/ko_KR/>      |

---

## How to Use

```bash
$ docker-compose up -d && docker-compose logs -f

# Go to http://[your domain]:[port]
# Default port # is 35000.
# You can get JSON: {"message":"Hello, this is go api server."}
```

---

## API Specifications

Refer [this page](../API_doc.md) for more information.

---

## References

### Go

- <https://hub.docker.com/_/golang>
- <https://soyoung-new-challenge.tistory.com/113>
- Unit Test
  - <http://golang.site/go/article/115-Go-유닛-테스트>
  - <https://lejewk.github.io/go-test/>
- CORS
  - <https://github.com/rs/cors>
- Live Reloading
  - <https://techinscribed.com/5-ways-to-live-reloading-go-applications/>
- JSON
  - <http://golang.site/go/article/104-JSON-사용>
  - <https://golang.org/pkg/encoding/json/>
  - <https://jeonghwan-kim.github.io/dev/2019/01/18/go-encoding-json.html>

### Gin

- <https://github.com/gin-gonic/gin#manually>
- <https://velog.io/@soosungp33/Golang으로-웹서버-만들기>
- <https://gin-gonic.com/ko-kr/docs/features/>
- <https://lv99.tistory.com/19>
- <https://bourbonkk.tistory.com/59>
- <https://riptutorial.com/ko/go/example/29299/gin을-사용한-restfull-프로젝트-api#undefined>
- Graceful shutdown
  - <https://github.com/gin-gonic/gin#graceful-shutdown-or-restart>
- Logging
  - <https://gin-gonic.com/ko-kr/docs/examples/write-log/>
  - <https://github.com/gin-gonic/gin#how-to-write-log-file>
- Time
  - <https://jeonghwan-kim.github.io/dev/2019/01/14/go-time.html>

### database/sql MariaDB

- <http://golang.site/go/article/107-MySql>
- <https://golang.org/pkg/database/sql/>
- <https://pkg.go.dev/github.com/go-sql-driver/mysql?utm_source=gopls#readme-installation>
- <https://mariadb.com/ko/resources/blog/using-go-with-mariadb/>
- <http://jameskim79.ipdisk.co.kr:8000/apps/wordpress/mariadb-명령어-및-사용법/>
- Connection Pool
  - <https://www.popit.kr/golang-databasesql-패키지-삽질기-3편-커넥션-풀/>
- NullString Conversion errors
  - <https://findstar.pe.kr/2018/01/23/null-scan-for-mysql-struct-golang/>
  - <https://golang.org/pkg/database/sql/#NullString>
- Time conversion errors
  - <https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time>

### swagger

- <https://victorydntmd.tistory.com/341>
- <https://velog.io/@sa833591/Swagger-API-문서-자동화>
- Gin에서 Swagger 사용하기 <https://dejavuqa.tistory.com/330>
