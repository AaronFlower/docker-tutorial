## Usage
- 启动 
```
docker-compose up
# or
docker-compose up -d
```
- 测试
```
bash -c 'i=0; while [ $i -lt 20 ]; do http :9000; i=$[$i + 1]; done;'|grep llo
```
