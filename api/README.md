# webpractice - api

## Environment

| Name | Version | Remarks                            |
|:----:|:-------:|:----------------------------------:|
| Go   | 1.15.7  | Docker version                     |
| Gin  | 1.6.3   | <https://github.com/gin-gonic/gin> |
| Gorm | 1.20.11 | <https://gorm.io/ko_KR/>           |

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
- <https://velog.io/@soosungp33/Golang으로-웹서버-만들기>
- <https://lv99.tistory.com/19>
- <http://golang.site/go/article/107-MySql-사용---쿼리>
- <http://golang.site/go/article/115-Go-유닛-테스트>
- <https://lejewk.github.io/go-test/>
- <https://github.com/rs/cors>

### Gorm

- <https://gorm.io/ko_KR/docs/models.html>

### Gin

- <https://gin-gonic.com/ko-kr/docs/features/>

### swagger

- <https://victorydntmd.tistory.com/341>
- <https://velog.io/@sa833591/Swagger-API-문서-자동화>
- Gin에서 Swagger 사용하기 <https://dejavuqa.tistory.com/330>
