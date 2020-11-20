:redis:
## 备份

### RDB(快照)

Redis Database Backup

- 配置: save [seconds] [changes]

### AOF

Append Only File
类似mysql binlog

- 配置开关: appendonly yes
- fsync配置：  appendfsync
  - no: 交给os处理
  - always： 每条日志操作
  - everysec: 折中做法，每秒进行一次

- AOF rewrite

## 内存管理
做缓存时可以手动定期删除key或者设置`maxmemory`&`maxmemory-policy`。

### maxmemory

### 淘汰策略:maxmemory-policy
- volatile-lru
  设置有效期的key lru

- allkeys-lru
  所有key-lru

- volatile-random
  随机淘汰，只淘汰设置有有效期的key

- allkeys-random
  随机淘汰数据

- volatile-ttl
  淘汰剩余有效期最短的

## 事物/scripting

## pipeline

连续且无相关的命令

---
[文章](https://app.getpocket.com/read/2402341457)


# redis-mindmap

通过思维导图整理redis的重要知识点

#### 一、持久化

https://github.com/Weiwf/redis-mindmap/blob/master/pic/%E6%8C%81%E4%B9%85%E5%8C%96.png

#### 二、复制

https://github.com/Weiwf/redis-mindmap/blob/master/pic/%E5%A4%8D%E5%88%B6.png

#### 三、阻塞

https://github.com/Weiwf/redis-mindmap/blob/master/pic/%E9%98%BB%E5%A1%9E.png

#### 四、Redis内存

https://github.com/Weiwf/redis-mindmap/blob/master/pic/Redis%E5%86%85%E5%AD%98.png

#### 五、Redis内存优化

(https://github.com/Weiwf/redis-mindmap/blob/master/pic/redis%E5%86%85%E5%AD%98%E4%BC%98%E5%8C%96.png

#### 六、哨兵

https://github.com/Weiwf/redis-mindmap/blob/master/pic/%E5%93%A8%E5%85%B5.png

#### 七、缓存设计

https://github.com/Weiwf/redis-mindmap/blob/master/pic/%E7%BC%93%E5%AD%98%E8%AE%BE%E8%A%A1.png

