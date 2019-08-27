# web网络安全

## 跨域

> - 不同域[^Origin]之间的请求。localhost和127.0.0.1为不同域，域名和解析到的IP不同域。

### 同源策略[^ SOP]

> 跨域可以发出请求，而是不被浏览器解析。

#### 同源策略限制内容

1. cookie、localStorage、indexDB等存储内容。
2. 不允许跨域dom操作。
3. ajax操作。

#### 同源策略没有限制的内容

`<img src=>`,`<like href= />`,`<script src= ></script>`,`<iframe src=></iframe>`等有*src*的标签
### 主要跨域请求实现方式
#### jsonp
**原理 **：利用`<script src>`标签可以跨域[^jsonp]。
**要点**

> 1. 只可以使用get方法。
> 2. 返回js代码。一般格式func(ReturnData)。
> 3. jquery ajax方法封装了jsonp方法。
> 4. 有安全风险。不建议使用。

#### CORS(Cross-Origin Resource Sharing)
##### 实现
**通过在js协议头中添加以下一些即可**

1. `Access-Control-Allow-Origin:Orign`。Orign：*****或者**<origin>**
2. `Access-Control-Allow-Methods:POST,GET`

#### websocket

## 常见安全

### xss
### xsrf/csrf
### 钓鱼
### 劫持
### 控制台

## 引用及链接
[^jsonp]: [stackoverflow关于jsonp原理的实现。](https://stackoverflow.com/questions/3839966/can-anyone-explain-what-jsonp-is-in-layman-terms)

[^Origin]: An origin is defined as a combination of [URI scheme](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier), [host name](https://en.wikipedia.org/wiki/Hostname), and [port number](https://en.wikipedia.org/wiki/Port_(computer_networking))

[^ SOP]: [**same-origin policy**](https://en.wikipedia.org/wiki/Same-origin_policy)

[常见网络安全](https://segmentfault.com/a/1190000006672214#articleHeader10)

[跨域的方式](https://blog.csdn.net/superaistar/article/details/83618689)

[实现跨域请求的八种方式](https://blog.csdn.net/ligang2585116/article/details/73072868)

