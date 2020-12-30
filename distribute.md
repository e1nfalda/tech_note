# redis 进阶
<!-- :redis: -->

## master-slave replication
通过`SYNC`命令同步master内容到slave。

## sentinel
监控节点运行状态，主从切换。

## redis-cluster
去中心化，使用gossip在集群间同步路由表和集群拓补信息,每个节点都保存相关信息。

*客户端向任一节点发送请求，根据节点返回值做重定向(MOVE、ASK)操作*

## codis
请求先请求到proxy,proxy做sharing后转发到相应实例上。`sharding规则`(路由表/转发表/slot表)保存在集中化组建(zookeeper, 文件系统等).
*codis在redis源码patch slot代码实现slot功能。*


----
[redis 集群](../../../../articles/Tech/redis-cluster和Codis关于slot迁移原理对比 - 黑客画家的个人空间 - OSCHINA __ Reader View.pdf)
