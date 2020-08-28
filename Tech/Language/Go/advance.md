
### GC

golang使用的[三色标记法](http://legendtkl.com/2017/04/28/golang-gc/)

> 常用的gc方法。引用计数；标记-清除;节点复制;分代收集；

### unsafe

> 因为golang普通指针无法进行数学运算等复杂操作。需要通过unsafe.Pointer类型进行转换及操作。
>
> unitptr:builtin类型。指针单位。32位系统4字节，64位8字节。

### methods

- **Sizeof(XType) uintptr**。获取*XType*类型占用的字节数
- **Offsetof(XType) unitptr**。结构体成员在结构体中位移。XType为结构体member。
- **Alignof(xType) unitptr**。指当类型进行内存对齐时，它分配到的内存地址能整除 m。

### unsafe.Pointer

```go
type ArbitraryType int
type Pointer *ArbitraryType 
```

|   uintptr   | unsafe.Pointer    | AnyType |
| ---- | ---- | ---- |
| uintptr(unsafePointer()) | unsafe.Pointer(&variable) |\*(Type\*)(unsafe.Pointer(&i))|

### reflect（反射/自省）

> reflect把数据从内存中映射到**Value**/**Type** struct中，及相关操作的接口集合。
>
> 反射的流程是把



### golang unpack时不能传递给interface{}类型参数

```go
func f(v ...interface{}) {
	fmt.Println("Hello, playground", v[0], v[1])
}

func main() {
	a := strings.Split("a,b,c", ",")
	f(a...)
}

// cannot use a (type []string) as type []interface {} in argument to f

```

### CSP 模型。

> 并发划分方式[wiki](https://zh.wikipedia.org/wiki/%E5%B9%B6%E8%A1%8C%E7%BC%96%E7%A8%8B%E6%A8%A1%E5%9E%8B)。进程交互，问题分解两种。
>
> ==进程交互==： **共享内存**，**消息（同步/异步）**，**分布式内存**。
>
> ==问题分解==：**数据并行**，**任务并行**，**隐式并行**

异步消息，任务并行方式的并行编程模型。

[string to interface{}](https://stackoverflow.com/questions/27689058/convert-string-to-interface)

https://golang.org/doc/faq#convert_slice_of_interface
