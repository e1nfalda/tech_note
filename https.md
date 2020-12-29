# HTTPS
:Tech:Theory:HTTPS:

![流程图](https://raw.githubusercontent.com/e1nfalda/IAaFaJdFLzSk/ignore/uPic/image-20200922142313775.png)

### 加密

分为`对称加密`和`非对称加密`。`对称加密`是加密和解密用同样密钥。`非对称加密`即用不同密钥。

#### 非对称加密

- RSA：算法实现简单，诞生于 1977 年，历史悠久，经过了长时间的破解测试，安全性高。缺点就是需要比较大的素数（目前常用的是 2048 位）来保证安全强度，很消耗 CPU 运算资源。RSA 是目前唯一一个既能用于密钥交换又能用于证书签名的算法。
- ECDHE：使用椭圆曲线（ECC）的 DH 算法，优点是能用较小的素数（256 位）实现 RSA 相同的安全等级。缺点是算法实现复杂，用于密钥交换的历史不长，没有经过长时间的安全攻击测试。

#### 对称加密

分`流失加密`和`分组加密`。

##### 流式加密

主要有`RC4`,不安全，不推荐使用。

`ChaCha20`：目前已经被 android 和 chrome 采用，也编译进了 google 的开源 openssl 分支---boring ssl，并且 nginx 1.7.4 也支持编译 boringssl。

#### 分组加密

分组加密以前常用的模式是 `AES-CBC`，但是 CBC 已经被证明容易遭受 BEAST 和 LUCKY13 攻击。目前建议使用的分组加密模式是 `AES-GCM`，不过它的缺点是计算量大，性能和电量消耗都比较高，不适用于移动电话和平板电脑。

### 2. 数字签名（digital signature）

数字签名可以认为是一个证书的防伪标签，目前使用最广泛的 SHA-RSA 数字签名的制作和验证过程如下：

1. 数字签名的签发。首先是使用哈希函数对证书数据哈希，生成消息摘要，然后使用 CA 自己的私钥对证书内容和消息摘要进行加密。
2. 数字签名的校验。使用 CA 的公钥解密签名，然后使用相同的签名函数对证书内容进行签名并和服务端的数字签名里的签名内容进行比较，如果相同就认为校验成功。

### 流程

1. 服务建立连接。server发送证书。
2. sender验证证书。生成random_key.采用证书上public key加密random key 发送server。

### Session cache

#### Session Identifier

Session Identifier（会话标识符），是 TLS 握手中生成的 Session ID。服务端可以将 Session ID 协商后的信息存起来，浏览器也可以保存 Session ID，并在后续的 ClientHello 握手中带上它，如果服务端能找到与之匹配的信息，就可以完成一次快速握手。

#### session ticket

一个会话ticket是一个加密的数据blob，其中包含需要重用的TLS连接信息，如会话key等，它一般是使用ticket key加密，因为ticket key服务器端也知道，在初始握手中服务器发送一个会话ticket到客户端，存储到客户端本地，当重用会话时，客户端发送会话ticket到服务器，服务器解密然后重用会话。

--------

[百度 https实践](https://developer.baidu.com/resources/online/doc/security/https-pratice-1.html)
