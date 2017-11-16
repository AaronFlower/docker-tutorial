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



### Hello world 进阶

有个问题，就是当我们的 Image 启动之后，当源文件变更了之后，我们的网站并没有实时刷新。怎样让我们的服务能够实时刷新那？我们用 `docker command --help` 来查看下 run 命令有那些选项。
```powershell
$ docker run --help                                                                                                                                                               ‹ruby-2.2.4›

Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container

Options:
      --add-host list                  Add a custom host-to-IP mapping (host:ip)
  -a, --attach list                    Attach to STDIN, STDOUT or STDERR
      --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
      --blkio-weight-device list       Block IO weight (relative device weight) (default [])
      --cap-add list                   Add Linux capabilities
      --cap-drop list                  Drop Linux capabilities
      --cgroup-parent string           Optional parent cgroup for the container
      --cidfile string                 Write the container ID to the file
      --cpu-period int                 Limit CPU CFS (Completely Fair Scheduler) period
      --cpu-quota int                  Limit CPU CFS (Completely Fair Scheduler) quota
      --cpu-rt-period int              Limit CPU real-time period in microseconds
      --cpu-rt-runtime int             Limit CPU real-time runtime in microseconds
  -c, --cpu-shares int                 CPU shares (relative weight)
      --cpus decimal                   Number of CPUs
      --cpuset-cpus string             CPUs in which to allow execution (0-3, 0,1)
      --cpuset-mems string             MEMs in which to allow execution (0-3, 0,1)
  -d, --detach                         Run container in background and print container ID
      --detach-keys string             Override the key sequence for detaching a container
      --device list                    Add a host device to the container
      --device-cgroup-rule list        Add a rule to the cgroup allowed devices list
      --device-read-bps list           Limit read rate (bytes per second) from a device (default [])
      --device-read-iops list          Limit read rate (IO per second) from a device (default [])
      --device-write-bps list          Limit write rate (bytes per second) to a device (default [])
      --device-write-iops list         Limit write rate (IO per second) to a device (default [])
      --disable-content-trust          Skip image verification (default true)
      --dns list                       Set custom DNS servers
      --dns-option list                Set DNS options
      --dns-search list                Set custom DNS search domains
      --entrypoint string              Overwrite the default ENTRYPOINT of the image
  -e, --env list                       Set environment variables
      --env-file list                  Read in a file of environment variables
      --expose list                    Expose a port or a range of ports
      --group-add list                 Add additional groups to join
      --health-cmd string              Command to run to check health
      --health-interval duration       Time between running the check (ms|s|m|h) (default 0s)
      --health-retries int             Consecutive failures needed to report unhealthy
      --health-start-period duration   Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
      --health-timeout duration        Maximum time to allow one check to run (ms|s|m|h) (default 0s)
      --help                           Print usage
  -h, --hostname string                Container host name
      --init                           Run an init inside the container that forwards signals and reaps processes
  -i, --interactive                    Keep STDIN open even if not attached
      --ip string                      IPv4 address (e.g., 172.30.100.104)
      --ip6 string                     IPv6 address (e.g., 2001:db8::33)
      --ipc string                     IPC namespace to use
      --isolation string               Container isolation technology
      --kernel-memory bytes            Kernel memory limit
  -l, --label list                     Set meta data on a container
      --label-file list                Read in a line delimited file of labels
      --link list                      Add link to another container
      --link-local-ip list             Container IPv4/IPv6 link-local addresses
      --log-driver string              Logging driver for the container
      --log-opt list                   Log driver options
      --mac-address string             Container MAC address (e.g., 92:d0:c6:0a:29:33)
  -m, --memory bytes                   Memory limit
      --memory-reservation bytes       Memory soft limit
      --memory-swap bytes              Swap limit equal to memory plus swap: '-1' to enable unlimited swap
      --memory-swappiness int          Tune container memory swappiness (0 to 100) (default -1)
      --mount mount                    Attach a filesystem mount to the container
      --name string                    Assign a name to the container
      --network string                 Connect a container to a network (default "default")
      --network-alias list             Add network-scoped alias for the container
      --pids-limit int                 Tune container pids limit (set -1 for unlimited)
      --privileged                     Give extended privileges to this container
  -p, --publish list                   Publish a container's port(s) to the host
  -P, --publish-all                    Publish all exposed ports to random ports
      --read-only                      Mount the container's root filesystem as read only
      --restart string                 Restart policy to apply when a container exits (default "no")
      --rm                             Automatically remove the container when it exits
      --stop-timeout int               Timeout (in seconds) to stop a container
      --storage-opt list               Storage driver options for the container
      --sysctl map                     Sysctl options (default map[])
      --tmpfs list                     Mount a tmpfs directory
  -t, --tty                            Allocate a pseudo-TTY
      --ulimit ulimit                  Ulimit options (default [])
  -u, --user string                    Username or UID (format: <name|uid>[:<group|gid>])
      --userns string                  User namespace to use
      --uts string                     UTS namespace to use
  -v, --volume list                    Bind mount a volume
      --volume-driver string           Optional volume driver for the container
      --volumes-from list              Mount volumes from the specified container(s)
  -w, --workdir string                 Working directory inside the container
```
我们可以看到 `-v, --volume list Bind mount a volume` 选项可以将本地目录挂载的到窗口的 volume 上，我们加一个选项就可以做实时刷新了。

```powershell
$ docker run -v ~/learning/git_repos/docker-tutorial/src:/var/www/html -p 8086:80 hello-world
```
现在我们就可刷新了。

注意： run 时的 options 选项会覆盖 Dockerfile 的配置，所以在运行进会以 run options 会准，但是并不会改变 Image 。在开developemnt 时可以用 options, 在 production 时，要记得重新 build.
