## Load Balance
### Get Started
先用 `docker-composer` 启动两个服务。

- docker-compose.yml
```yml
version: '3'

services:
    website1:
        image: nginx:1.15-alpine
        volumes:
            - ./home1:/usr/share/nginx/html
        ports:
            - 9001:80

    website2:
        image: nginx:1.15-alpine
        volumes:
            - ./home2:/usr/share/nginx/html
        ports:
            - 9002:80
```
- 启动
```
docker-compose up
```
- 访问
```
http :8081
http :8082
```

- 输出
```
04-load-balance master ✗ 46d ◒ ➜ docker-compose up
Creating network "04-load-balance_default" with the default driver
Creating 04-load-balance_website2_1 ... done
Creating 04-load-balance_website1_1 ... done
Attaching to 04-load-balance_website1_1, 04-load-balance_website2_1
website1_1  | 172.18.0.1 - - [12/Dec/2018:03:52:01 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2_1  | 172.18.0.1 - - [12/Dec/2018:03:52:04 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1_1  | 172.18.0.1 - - [12/Dec/2018:03:52:09 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2_1  | 172.18.0.1 - - [12/Dec/2018:03:52:11 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2_1  | 172.18.0.1 - - [12/Dec/2018:03:53:07 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1_1  | 172.18.0.1 - - [12/Dec/2018:03:53:09 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
```

### 使用 nginx 代理转发

