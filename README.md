## Docker Tutorial

### 1.安装  Docker

[Mac 下载安装](/Users/easonzhan/learning/git_repos/docker-tutorial)

```powershell
$ docker version                                                        
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
$ tree                                                                  
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
$ docker build -t hello-world . # 在当前目录下 build 会读取我们 Dockerfile  
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
$ docker run -p 8086:80 hello-world    
AH00558: apache2: Could not reliably determine the server's fully qualified domain name, using 172.17.0.2. Set the 'ServerName' directive globally to suppress this message
AH00558: apache2: Could not reliably determine the server's fully qualified domain name, using 172.17.0.2. Set the 'ServerName' directive globally to suppress this message
[Thu Nov 16 12:30:20.825095 2017] [mpm_prefork:notice] [pid 1] AH00163: Apache/2.4.10 (Debian) PHP/7.1.11 configured -- resuming normal operations
[Thu Nov 16 12:30:20.825289 2017] [core:notice] [pid 1] AH00094: Command line: 'apache2 -D FOREGROUND'
172.17.0.1 - - [16/Nov/2017:12:30:42 +0000] "GET / HTTP/1.1" 200 242 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
```

通过浏览器访问 `localhost:8086` 就可以看到 `hello world!` 了。

### Docker Images 

Docker 中一个最重要的概念就是 Image 镜像了。上面的我例子中我们 build 了一个镜像。下面是关于 Images 的一些命令。

#### 查看
```
➜ docker help images

Usage:  docker images [OPTIONS] [REPOSITORY[:TAG]]

List images

Options:
  -a, --all             Show all images (default hides intermediate images)
      --digests         Show digests
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print images using a Go template
      --no-trunc        Don't truncate output
  -q, --quiet           Only show numeric IDs
```
-  查看所有的镜像
```
$ docker images -a
```
- 查看 dangling 的镜像
```
$ docker images -f dangling=true
```

#### 删除镜像



Docker provides a single command that will clean up any resources — images, containers, volumes, and networks — that are dangling (not associated with a container):

```
docker system prune
```

### Docker 容器
通过上面的 `docker run ` 我们就启动了一个容器。

下面是窗口的一些命令。

```
docker help container

Usage:  docker container COMMAND

Manage containers

Options:


Commands:
  attach      Attach local standard input, output, and error streams to a running container
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  inspect     Display detailed information on one or more containers
  kill        Kill one or more running containers
  logs        Fetch the logs of a container
  ls          List containers
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  prune       Remove all stopped containers
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  run         Run a command in a new container
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  wait        Block until one or more containers stop, then print their exit codes

Run 'docker container COMMAND --help' for more information on a command.
```

- 查看
```
docker container ls
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                  NAMES
e8b4ef599dc4        hello-world         "docker-php-entrypoi…"   23 seconds ago      Up 25 seconds       0.0.0.0:8083->80/tcp   jovial_stallman
```

### Hello world 进阶

有个问题，就是当我们的 Image 启动之后，当源文件变更了之后，我们的网站并没有实时刷新。怎样让我们的服务能够实时刷新那？我们用 `docker command --help` 来查看下 run 命令有那些选项。
```powershell
$ docker run --help                                                       

Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container

Options:

  -v, --volume list                    Bind mount a volume
      --volume-driver string           Optional volume driver for the container
      --volumes-from list              Mount volumes from the specified container(s)
  -w, --workdir string                 Working directory inside the container
```
我们可以看到 `-v, --volume list Bind mount a volume` 选项可以将本地目录挂载的到容器的 volume 上，我们加一个选项就可以做实时刷新了。

```powershell
$ docker run -v ~/learning/git_repos/docker-tutorial/src:/var/www/html -p 8086:80 hello-world
```
现在我们就可刷新了。

注意： 执行 `run ` 命令中的配置项会覆盖 Dockerfile 文件中的配置，所以在运行时会以 run options 会准，但是并不会改变 Image 。在开发环境中 (developemnt) 可以用 options, 在生产环境中(production) 要记得重新 build.
