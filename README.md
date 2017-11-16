## Docker Tutorial

### 1.安装  Docker

[Mac 下载安装](/Users/easonzhan/learning/git_repos/docker-tutorial)

```powershell
$ docker version                                                                  ‹ruby-2.2.4›
Client:
 Version:      17.06.0-ce
 API version:  1.30
 Go version:   go1.8.3
 Git commit:   02c1d87
 Built:        Fri Jun 23 21:31:53 2017
 OS/Arch:      darwin/amd64
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
```

### 2. Hello world

Docker 用 Dockerfile 来管理 app, 在这个例子中我们用 php + apache 启动一个服务器。这个简单的例子目录结构为：

```powershell
$ tree                                                                            ‹ruby-2.2.4›
.
├── Dockerfile
├── Readme.md
└── src
    └── index.php

1 directory, 3 files
```

因为我们服务器依赖 php, 我们用 `docker search php`到 docker hub 找一下 app. 当然也可直接到 [Docker Hub](https://hub.docker.com/) 搜索。找到一个 Official 的 Image 来用。

我们用下面这个版本。

- [`7.1.11-apache-jessie`, `7.1-apache-jessie`, `7-apache-jessie`, `apache-jessie`, `7.1.11-apache`, `7.1-apache`, `7-apache`, `apache` (*7.1/jessie/apache/Dockerfile*)](https://github.com/docker-library/php/blob/bfe27759103fa6050601060165409b5b3be06395/7.1/jessie/apache/Dockerfile)

编写 `Dockerfile`.

```dockerfile
FROM php:7.1-apache
copy src/ /var/www/html

# Expose apache 80 # docker 里的 apache 启动的是 80 端口。
EXPOSE 80
```

build 我们的 image.

```powershell
$ docker build -t hello-world . # 在当前目录下 build 会读取我们 Dockerfile                                                                                                                                                  ‹ruby-2.2.4›
Sending build context to Docker daemon  5.632kB
Step 1/2 : FROM php:7.1-apache
7.1-apache: Pulling from library/php
85b1f47fba49: Pull complete
d8204bc92725: Pull complete
92fc16bb18e4: Pull complete
31098e61b2ae: Pull complete
f6ae64bfd33d: Pull complete
003c1818b354: Pull complete
a6fd4aeb32ad: Pull complete
a094df7cedc1: Pull complete
af0f77e732e0: Pull complete
1513b36e0001: Pull complete
f0e4a4e2be44: Pull complete
b050de8f5d3e: Pull complete
f595ab2f751a: Pull complete
Digest: sha256:19c0c242f90da77a0f931bab9357fb2b497f040f22aee362275ba53ffa4be4e8
Status: Downloaded newer image for php:7.1-apache
 ---> cb6a5015ad72
Step 2/2 : COPY src/ /var/www/html
 ---> fd789317eb26
Removing intermediate container a6dcf9b63ac2
Successfully built fd789317eb26
Successfully tagged hello-world:latest
```

build 之后就可运行我们的 docker Image了。因为我本地的机器的 `80` 已经被占用了，所以用 `8086`端口来接收 docker expose 的 `80`  端口。

```powershell
$ docker run -p 8086:80 hello-world                                               ‹ruby-2.2.4›
AH00558: apache2: Could not reliably determine the server's fully qualified domain name, using 172.17.0.2. Set the 'ServerName' directive globally to suppress this message
AH00558: apache2: Could not reliably determine the server's fully qualified domain name, using 172.17.0.2. Set the 'ServerName' directive globally to suppress this message
[Thu Nov 16 12:30:20.825095 2017] [mpm_prefork:notice] [pid 1] AH00163: Apache/2.4.10 (Debian) PHP/7.1.11 configured -- resuming normal operations
[Thu Nov 16 12:30:20.825289 2017] [core:notice] [pid 1] AH00094: Command line: 'apache2 -D FOREGROUND'
172.17.0.1 - - [16/Nov/2017:12:30:42 +0000] "GET / HTTP/1.1" 200 242 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
```

通过浏览器访问 `localhost:8086` 就可以看到 `hello world!` 了。