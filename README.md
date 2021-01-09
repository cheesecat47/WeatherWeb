# webpractice

> Project for studying Kafka, Rest api, etc.

---

## Environment

| Name | Version | Link |
|:-:|:-:|:-:|
| Docker | 20.10.0 | <https://docs.docker.com/engine/install/> |
| Docker-compose | 1.27.4 | <https://docs.docker.com/compose/install/> |

* tmux & [tmuxinator](https://github.com/tmuxinator/tmuxinator) (additional)

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

1. Else

---

## Deploy Ports

| Dir | Name | Port # |
|:-:|:-:|:-:|
| kafka | zookeeper | 31010 |
| kafka | kafka | 31020 |
| nginx | nginx | 32000 |
| dbserver | mariadb | 33010 |
| dbserver | adminer | 33020 |

---

## Contributors' info
  
An Taegeon - <https://github.com/atg0831>  
Han Jinkyu - <https://github.com/jinkyuhan>  
Shin Juyong – <https://github.com/cheesecat47>  

[https://github.com/cheesecat47/webpractice](https://github.com/cheesecat47/webpractice)  

## References

<https://hanseokhyeon.tistory.com/entry/tmuxinator-custom-layout-%EB%A7%8C%EB%93%A4%EA%B8%B0>
