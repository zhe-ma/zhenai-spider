### 珍爱网爬虫学习

1. Docker安装ElasticSearch。
`$ docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2`

2. 运行容器。
`$ docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2`

3. 修改配置文件conf/config.yaml，配置ElasticSearch的地址。
`host: http://10.196.102.145:9200`

4. 运行爬虫
`$ cd zhenai-spider`
`$ go run main.go`

5. 网页展示
`$ cd frontend`
`$ go run starter.go`

6. 浏览器输入`http://localhost:9571/`。

7. 搜索测试：
`Age:30`
`Height:>166`
`Sex:male Height:([166 TO 168])`