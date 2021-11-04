### redis
>Redis是一个开源的、使用C语言编写的、支持网络交互的、可基于内存也可持久化的Key-Value数据库。Redis 通常被称为`数据结构服务器`，因为值（value）可以是字符串(String)、哈希(Hash)、列表(list)、集合(sets)和有序集合(sorted sets)等类型。


+ 相关资源

[1. Redis 官网](https://redis.io/)
[2. 源码地址](https://github.com/redis/redis)
[3. Redis 在线测试](http://try.redis.io/)
[4. Redis 命令参考](http://doc.redisfans.com/)

+ 优势
1. 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s 。
2. 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。
3. 原子 – Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来。
4. 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。

+ ubuntu安装 apt-get 命令
```
sudo apt-get update
sudo apt-get install redis-server

redis-server    #启动redis

redis-cli # 查看是否启动

chccc@chccc:~$ redis-cli
127.0.0.1:6379> ping
PONG
127.0.0.1:6379> 
```
+ REDIS配置
1. 获取配置
语法： CONFIG GET CONFIG_SETTING_NAME
```
CONFIG GET *

CONFIG GET loglevel
```
2. 设置配置
语法：CONFIG SET CONFIG_SETTING_NAME NEW_CONFIG_VALUE
```
CONFIG SET loglevel "notice"
```

