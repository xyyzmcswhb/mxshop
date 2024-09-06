# mxshop

#### 介绍

本项目是一个网上生鲜超市，可以完成加购物车，收藏，下单，结算等购物的相关功能，本项目采用手机号动态验证码注册，支付宝支付等功能，为大家提供了一个完美的大型电商超市。

电商系统 - 分层式微服务架构，服务之间互相解耦。

service层向内提供grpc通讯，web层向外提供http通讯。

业务逻辑在web层开发，底层交互在service层开发。

所有服务集成nacos配置中心、consul注册服务发现、健康检查、负载均衡。

redis,rocketmq中间件，mysql存储，elasticsearch全文搜索。

#### 软件架构

软件架构说明


#### 安装教程

本人环境：macos m1 silicon
测试工具：postman,apifox,yapi皆可

##### 1、安装Nacos
https://www.runoob.com/docker/macos-docker-install.html
或者直接使用docker安装

```shell
# intel 处理器安装
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=128m -e JVM_XMX=128m -e JVM_XMN=128m -p 8848:8848  -p 9848:9848 -p 9555:9555 -d nacos-server:lates


#  Apple Silicon 处理器安装
docker run --name nacos-standalone -e MODE=standalone  -p 8848:8848 -p 9848:9848 -p 9555:9555  -e JVM_XMS=128m -e JVM_XMX=128m -e JVM_XMN=128m -d nacos/nacos-server:v2.1.0-slim
```

##### 2、安装Consul
macos下安装
https://cloud.tencent.com/developer/article/1890248

##### 3、安装Elasticsearch
https://blog.csdn.net/the_shy_faker/article/details/128520129
推荐使用docker安装
```shell
#新建es的config配置文件夹
mkdir -p /data/elasticsearch/config
#新建es的data目录
mkdir -p /data/elasticsearch/data
#新建es的plugins目录
mkdir -p /data/elasticsearch/plugins
#给目录设置权限
chmod 777 -R /data/elasticsearch
#写入配置到elasticsearch.yml中， 下面的 > 表示覆盖的方式写入， >>表示追加
echo "http.host: 0.0.0.0" >> /data/elasticsearch/config/elasticsearch.yml
#安装es
docker run --name elasticsearch -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e ES_JAVA_OPTS="-Xms128m -Xmx256m" \
-v /data/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config
-v /data/elasticsearch/data:/usr/share/elasticsearch/data \
-v /data/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
-d elasticsearch:7.10.1
```

es可视化工具kibana可选择性安装,postman也可

##### 4.安装rocketmq
推荐使用docker安装
https://juejin.cn/post/7160983695044804644


#### 配置示例
web层
```json
{
  "name": "user-web",
  "host": "127.0.0.1",
  "tags" : ["hb","user","fighting"],
  "port": 8081,
  "user_srv": {
    "host": "127.0.0.1",
    "port": "50050",
    "name": "user-srv"
  },
  "jwt": {
    "key": "********"
  },
  "sms": {
    "key": "**********",
    "secret": "********"
  },
  "redis": {
    "host": "127.0.0.1",
    "port": 6379,
    "expire": 300
  },
  "consul": {
    "host": "127.0.0.1",
    "port": 8500
  }
}
```
srvs层

```json
{
  "name": "user-srv",
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "123456hb",
    "db": "mxshop_user_srv"
  },
  "consul": {
    "host": "127.0.0.1",
    "port": 8500
  }
}
```


