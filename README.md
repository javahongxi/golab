# golab
```
git clone https://github.com/javahongxi/golab.git
go get ./...
```

### About GoLab
- lang: Go语法学习与经典示例
- crawler: 爬虫项目单任务版及并发版
- crawler_distributed: 爬虫项目分布式版
- pipeline: 搭建并行处理管道

### 爬虫项目并发版演示
- 启动相亲网站 [mockserver](https://github.com/javahongxi/mockserver)
```shell
go run mockserver/main.go
```
- 启动爬虫
```shell
go run crawler/main.go
```

### 爬虫项目分布式版演示
- 启动相亲网站 [mockserver](https://github.com/javahongxi/mockserver)
```shell
go run mockserver/main.go
```
- 下载ES并本地启动
```shell
bin/elasticsearch
```
- 启动存储服务(数据存储到ES)
```shell
cd crawler_distributed/persist/server
go run itemsaver.go --port=9090
```
- 启动多个worker服务(爬取数据)
```shell
cd crawler_distributed/worker/server
go run worker.go --port=9091
go run worker.go --port=9092
```
- 启动爬虫
```shell
cd crawler_distributed
go run main.go --itemsaver_host=:9090 --worker_hosts=:9091,:9092
```

### Go Projects
- https://github.com/moby/moby
- https://github.com/docker/docker-ce
- https://github.com/kubernetes/kubernetes
- https://github.com/etcd-io/etcd
- https://github.com/gin-gonic/gin
- https://github.com/hashicorp/consul
- https://github.com/micro/go-micro
- https://github.com/nsqio/nsq
- https://github.com/elastic/beats
- https://github.com/pingcap/tidb
- https://github.com/CodisLabs/codis
- https://github.com/baidu/bfe
- https://github.com/caddyserver/caddy
- https://github.com/cockroachdb/cockroach

&copy; [hongxi.org](http://hongxi.org) | [go.hongxi.org](http://go.hongxi.org)
