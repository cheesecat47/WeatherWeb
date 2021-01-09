# WeatherWeb - db

## Environment

| Name | Version | Link |
|:-:|:-:|:-:|
| MariaDB | 10.5.8 | <https://hub.docker.com/_/mariadb> |

---

## How to Use

```bash
$ docker-compose up -d && docker-compose logs -f
```

```bash
$ mysql -u root -p webpracticedb
```

Down the components
```bash
$ docker-compose down
$ rm -rf ../data/mariadb
```

---

## References

<https://mariadb.org/>  
<https://mariadb.com/kb/ko/mariadb/>  
<https://engkimbs.tistory.com/931>  
<https://hub.docker.com/_/mariadb>  
<https://happygrammer.github.io/docker/mariadb/>  
<https://kamang-it.tistory.com/entry/DockerMysqlInitSql>  
<https://judo0179.tistory.com/69>  
