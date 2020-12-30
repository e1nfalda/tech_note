# 网络工具
<!-- :Tech:tools:network: -->

---

### 命令
#### route：macOS设置系统路由的命令。
#### netstat: 显示网路状态。

>  The netstat command symbolically displays the contents of various network-related data structures.  There are a number of output formats,
>      depending on the options for the information presented.  The first form of the command displays a list of active sockets for each proto-
>      col.  The second form presents the contents of one of the other network data structures according to the option selected. Using the third
>      form, with a wait interval specified, netstat will continuously display the information regarding packet traffic on the configured net-
>      work interfaces.  The fourth form displays statistics for the specified protocol or address family. If a wait interval is specified, the
>      protocol information over the last interval seconds will be displayed.  The fifth form displays per-interface statistics for the speci-
>      fied protocol or address family.  The sixth form displays mbuf(9) statistics.  The seventh form displays routing table for the specified
>      address family.  The eighth form displays routing statistics

#### pfctl: 控制pf（package filter）和nat（network address translation）设备。

### miscellanies

#### `netstat -nr`打印路由表 

#### macOS interface 作用

> `en0` at one point "ethernet", now is WiFi (and I have no idea what extra `en1` or `en2` are used for).
>
> `fw0` is the FireWire network interface.
>
> `stf0` is an [IPv6 to IPv4 tunnel interface](https://www.freebsd.org/cgi/man.cgi?gif(4)) to support [the transition](http://en.wikipedia.org/wiki/6to4) from IPv4 to the IPv6 standard.
>
> `gif0` is a more [generic tunneling interface](https://www.freebsd.org/cgi/man.cgi?gif(4)) [46]-to-[46].
>
> `awdl0` is [Apple Wireless Direct Link](https://stackoverflow.com/questions/19587701/what-is-awdl-apple-wireless-direct-link-and-how-does-it-work)
>
> `p2p0` is related to AWDL features. Either as an old version, or virtual interface with different semantics than `awdl`.



### gateway

TCP/IP协议中网络设备分两类，主机（host）和网关（gateway）。

> The term gateway can also loosely refer to a computer or computer program configured to perform the tasks of a gateway, such as a [default gateway](https://en.wikipedia.org/wiki/Default_gateway) or [router](https://en.wikipedia.org/wiki/Router_(computing)).
>
> In enterprise networks, a network gateway usually also acts as a [proxy server](https://en.wikipedia.org/wiki/Proxy_server) and a [firewall](https://en.wikipedia.org/wiki/Firewall_(computing)).

使用不同通信协议的设备之间数据传输的设备。

**Default gateway**：If you add a routing entry to route the traffic destined to reach any destination pointing a next hop address , the next hop is your default gateway.
