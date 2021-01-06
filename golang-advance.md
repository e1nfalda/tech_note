# Golang 进阶
<!-- :Tech:Language: -->

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

| uintptr                  | unsafe.Pointer            | AnyType                        |
| -----------------------  | ----                    | ----                           |
| uintptr(unsafePointer()) | unsafe.Pointer(&variable) | \*(Type\*)(unsafe.Pointer(&i)) |

### reflect（反射/自省）

> reflect把数据从内存中映射到**Value**/**Type** struct中，及相关操作的接口集合。
>
> 反射的流程是把

```go
// demo
if reflect.ValueOf(e).Field(i).Kind() != reflect.Struct {
    fmt.Println(reflect.ValueOf(e).Field(i))
}
```

### golang unpack时不能传递给interface{}类型参数

```go
func f(v ...interface{}) {
	fmt.Println("Hello, playground", v[0], v[1])
}

func main() {
	a := strings.Split("a,b,c", ",")
	f(a...)
}

// ❌ cannot use a (type []string) as type []interface {} in argument to f

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



### new vs make

两者都是在堆上分配。

`new`: `new(Type) *Type`分配内存初始化(置0)变量并返回变量指针。`Type`类型 int，type定义struct等。

`make`: make(Type, args...) Type  初始化变量返回变量。针对类型 slice, map,channel。因为类型本身是引用类型，没必要返回指针。

````go
// ❌ 编辑错误
...
var i *int
*int = 10 

// ✅ new初始化int内存，并返回
...
var i *int = new(int)
*i = 10 

// ✅ new初始化struct，并返回对应指针
type user struct {
	name string
}
func main() {
	u := new(user)
	u.name = "wong"
	fmt.Println("Hello, playground", u) // 
}

````
### go install and uninstall
`go get IMPORT_PATH` <==> `go clean -i IMPORT_PATH`

