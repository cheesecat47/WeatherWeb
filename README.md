# webpractice

> Project for studying Kafka, Rest api, etc.

---

## Environment

| Name | Version | Link |
|:-:|:-:|:-:|
| Docker | 20.10.0 | <https://docs.docker.com/engine/install/> |
| Docker-compose | 1.27.4 | <https://docs.docker.com/compose/install/> |
| tmux & tmuxinator  (optional) |  | <https://github.com/tmuxinator/tmuxinator> |


---

## How to Use

```bash
git clone https://github.com/cheesecat47/webpractice.git
cd webpractice
```

1. If you installed tmux & tmuxinator, run this.  

    ```bash
    $ tmuxinator start project
    ```

1. If you don't use tmux, run this.  

    ```bash
    $ sudo docker-compose up -d && sudo docker-compose logs -f
    ```

---

## Deploy Ports

| Dir      | Name       | Port # |
|:--------:|:----------:|:------:|
| kafka    | kafka_zk   | 31010  |
| kafka    | kafka      | 31020  |
| nginx    | nginx      | 32000  |
| dbserver | db_maria   | 33010  |
| dbserver | db_adminer | 33020  |
| react    | react      | 34000  |
| api      | api        | 35000  |

---

## Contributors' info
  
An Taegeon - <https://github.com/atg0831>  
Han Jinkyu - <https://github.com/jinkyuhan>  
Shin Juyong – <https://github.com/cheesecat47>  

[https://github.com/cheesecat47/webpractice](https://github.com/cheesecat47/webpractice)  

## References

<https://hanseokhyeon.tistory.com/entry/tmuxinator-custom-layout-%EB%A7%8C%EB%93%A4%EA%B8%B0>
