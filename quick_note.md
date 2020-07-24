## Go

### CSP 模型。

> 并发划分方式[wiki](https://zh.wikipedia.org/wiki/%E5%B9%B6%E8%A1%8C%E7%BC%96%E7%A8%8B%E6%A8%A1%E5%9E%8B)。进程交互，问题分解两种。
>
> ==进程交互==： **共享内存**，**消息（同步/异步）**，**分布式内存**。
>
> ==问题分解==：**数据并行**，**任务并行**，**隐式并行**

异步消息，任务并行方式的并行编程模型。



os.Args,

 os.Open(),op.Create, os.OpenFile(), os.Close():

sync.WaitGroup  wg.Add, wg.Wait

sync.Lock, sync.Unlock

http.



%t : Boolean

%T: typeof

%e, %f: float

### channel

make (chan val-type, buffer_size)

#### select

> `default`: 表示非阻塞。

```go
select {
    case var1 <- chan1:
    	...
    default:
    	...
}

```

#### 关闭

```go
if reflect.ValueOf(e).Field(i).Kind() != reflect.Struct {
    fmt.Println(reflect.ValueOf(e).Field(i))
}
```



==type aliase==



## javascript

According to the [ES5 spec](http://es5.github.com/), when doing [bitwise operations](http://es5.github.com/#x11.10), all operands are converted to [`ToInt32`](http://es5.github.com/#x9.5) (which first calls [`ToNumber`](http://es5.github.com/#x9.3). If the value is `NaN` or `Infinity`, it's converted to `0`).