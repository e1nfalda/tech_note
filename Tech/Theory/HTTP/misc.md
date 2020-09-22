#### `GET` vs `POST`
1. 缓存：get主动缓存。post不。
2. 幂等性。get是。post不一定。
3. get一次发送请求。post一般先发送*header*,服务器响应continue code（100），发送body。
4. 参数位置不同。
5. 编码。get：ascii。post无限制。

#### http connection

`hop-to-hop`  `end-to-end`, `short live connection`, `persistence connection`,`http piplining`

`tcp keep-alive` vs `http keep alive`:防火墙、负载均衡



----

[http keep alive](https://xie.infoq.cn/article/398b82c2b4300f928108ac605)

