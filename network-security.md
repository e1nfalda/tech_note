# web网络安全
:web-security:

## 同域&跨域

> - 不同域[^Origin]之间的请求。localhost和127.0.0.1为不同域，域名和解析到的IP不同域。

### 同源策略[^ SOP]

> - 跨域可以发出请求，但不被浏览器解析。

同源策略限制的内容如下：

- cookie、localStorage、indexDB等存储内容。
- 不允许跨域dom操作。
- ajax操作。

### 跨域方式

### document.domain

> If two windows (or frames) contain scripts that set domain to the same value, the same-origin policy is relaxed for these two windows, and each window can interact with the other. （维基百科）

#### jsonp

**原理 **：利用`<script .src>`标签可以跨域[^jsonp]cross-origin embedding。

> `<img .src=>`,`<like href= />`,`<script src= ></script>`,`<iframe src=></iframe>`等有*src*的标签(cross-origin embedding).

**要点**

> 1. 只可以使用get方法。
> 2. 返回js代码。一般格式func(ReturnData)。
> 3. jquery ajax方法封装了jsonp方法。
> 4. 有安全风险。不建议使用。

#### CORS(Cross-Origin Resource Sharing)

> The other technique for relaxing the same-origin policy is standardized under the name [Cross-Origin Resource Sharing](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing). This standard extends HTTP with a new Origin request header and a new Access-Control-Allow-Origin response header.[[8\]](https://en.wikipedia.org/wiki/Same-origin_policy#cite_note-8) It allows servers to use a header to explicitly list origins that may request a file or to use a wildcard and allow a file to be requested by any site. Browsers such as Firefox 3.5, Safari 4 and Internet Explorer 10 use this header to allow the cross-origin HTTP requests with XMLHttpRequest that would otherwise have been forbidden by the same-origin policy.(维基百科)

##### 实现

**通过在js协议头中添加以下一些即可**

1. `Access-Control-Allow-Origin:Orign`。Orign：*****或者**<origin>**
2. `Access-Control-Allow-Methods:POST,GET`

#### websocket

> websocket 建立连接无同源限制。

## 常见安全问题：

### xss

> 在浏览器窃取本地cookie等资源，模拟用户请求窃取用户数据，模拟用户请求等。

### xsrf/csrf（Cross Site Request Forgery, 跨站域请求伪造）

### 钓鱼

### 劫持

### 控制台

## 引用及链接

[^ SOP]: [**same-origin policy**](https://en.wikipedia.org/wiki/Same-origin_policy)

[常见网络安全](https://segmentfault.com/a/1190000006672214#articleHeader10)

[跨域的方式](https://blog.csdn.net/superaistar/article/details/83618689)

[实现跨域请求的八种方式](https://blog.csdn.net/ligang2585116/article/details/73072868)

[^jsonp]: [stackoverflow关于jsonp原理的实现。](https://stackoverflow.com/questions/3839966/can-anyone-explain-what-jsonp-is-in-layman-terms)
An origin is defined as a combination of [URI scheme](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier), [host name](https://en.wikipedia.org/wiki/Hostname), and [port number](https://en.wikipedia.org/wiki/Port_(computer_networking))
