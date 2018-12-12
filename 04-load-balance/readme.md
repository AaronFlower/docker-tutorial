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

# 你也可以在后台启动，进行守护，使用 -d 参数。
docker-compose up -d

# 查看 
docker-compose ps
           Name                    Command          State          Ports
--------------------------------------------------------------------------------
04-load-balance_website1_1   nginx -g daemon off;   Up      0.0.0.0:9001->80/tcp
04-load-balance_website2_1   nginx -g daemon off;   Up      0.0.0.0:9002->80/tcp
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
我们可再启用一个 nginx 作为转发服务器。来对布置在两台机器上的服务进行代理，来做下负载均衡。则其 `docker-compose.yml` 服务器的配置如下：
```yml
version: '3'

services:
    nginx:
        image: nginx:1.15-alpine
        volumes:
            - ./nginx.conf:/etc/nginx/conf.d/nginx.conf
        ports:
            - 9000:9000
        networks:
            - myapp
    website1:
        image: nginx:1.15-alpine
        volumes:
            - ./home1:/usr/share/nginx/html
        ports:
            - 80
        networks:
            - myapp
        container_name: website1
        restart: always
    website2:
        image: nginx:1.15-alpine
        volumes:
            - ./home2:/usr/share/nginx/html
        ports:
            - 80
        networks:
            - myapp
        container_name: website2
        restart: always
networks:
    myapp:
```

我们的 `nginx.conf` 配置文件如下：
```nginx

upstream myupload {
   server website1:80 weight=5;
   server website2:80 weight=1;
}

server {
    listen 9000;

    server_name localhost;

    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://myupload;
    }
}
```

测试如果如下：
```

04-load-balance master ✗ 3h31m △ ◒ ➜ bash -c 'i=0; while [ $i -lt 20 ]; do http :9000; i=$[$i + 1]; done;'|grep h1
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 2</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 2</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
        <h1>Home 2</h1>
        <h1>Home 1</h1>
        <h1>Home 1</h1>
```
nginx 转发记录如下：
```bash
04-load-balance master ✗ 3h28m △ ◒ ➜ docker-compose up
Creating network "04-load-balance_myapp" with the default driver
Creating website1                ... done
Creating website2                ... done
Creating 04-load-balance_nginx_1 ... done
Attaching to website2, 04-load-balance_nginx_1, website1
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:26:41 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:26:41 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
website1    | 172.28.0.3 - - [12/Dec/2018:07:26:44 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:26:44 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:26:45 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:26:45 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2    | 172.28.0.3 - - [12/Dec/2018:07:26:47 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:26:47 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:49 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:49 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:49 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:49 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:50 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2    | 172.28.0.3 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:51 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:52 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2    | 172.28.0.3 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:53 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website2    | 172.28.0.3 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:54 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:55 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:55 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
nginx_1     | 172.28.0.1 - - [12/Dec/2018:07:29:55 +0000] "GET / HTTP/1.1" 200 226 "-" "HTTPie/0.9.9" "-"
website1    | 172.28.0.3 - - [12/Dec/2018:07:29:55 +0000] "GET / HTTP/1.0" 200 226 "-" "HTTPie/0.9.9" "172.28.0.1"
```
