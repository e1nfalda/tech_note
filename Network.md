
# Network
>以太包size: 1500 bit*payload*+22bit *head* = 1522 （最初共1518）
>
>ip包：1480 + 20
>
>tcp包：1480 - 1420
>
>http包：HTTP/2 协议的一大改进， 就是压缩 HTTP 协议的头信息，使得一个 HTTP 请求可以放在一个 TCP 数据包里面

## 分层(tcp/ip 5层,OSI 7层)

![net_protocols](../.././.src/net_protocols.png)

### 链路层：实现了hop间的通信协议。PPP、ATM、Ethernet等

> Frame(数据帧)

### 网络层：IP，路由协议。ARP

> Packet(数据包)

### 传输层：实现tcp/udp。
  * [TCP](TCP)

