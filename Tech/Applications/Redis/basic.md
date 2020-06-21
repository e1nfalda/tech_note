#redis
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

