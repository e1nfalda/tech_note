#### `GET` vs `POST`
1. 缓存：get主动缓存。post不。
2. 幂等性。get是。post不一定。
3. get一次发送请求。post一般先发送*header*,服务器响应continue code（100），发送body。
4. 参数位置。
5. 编码。get：ascii.post无限制。