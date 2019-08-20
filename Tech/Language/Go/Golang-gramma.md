# Golang 语法（golang tour）

### var

1. **var** statment can be at package or function level;
2. 例：`var i,j int = 1, 2`

### TYPE CONVERSION

1. `T(v)`

### const

例： `const x = V`。

1. V  类型： bool、string、chacter、numeric values;

### for

1. *initial statement* and *post statment* can be omited.  

   ```go
   for [i := 1;] i < 10 [; i ++] {
       // todo
   } 
   ```

2. act as **while**

   ``` go
   for x < 100 {
       // todo
   }
   ```

3. forever Loop

   ```go
   for {
       // todo
   }
   ```

### if

1. can start with short statment；

```go
/* short statment x:= 10, 变量x的作用域仅在 if作用域内 */
if x := 10; x > y {
    // todo
}
```

2. 格式。

   ```go
   if Expresion {
       ...
   } else {
       ...
   }
   ```

### switch

1. short statment.

   ```go
   switch os := runtime.GOOS; os {
   case "linux":
     /* case 锁紧和switch同行 */
     ...
     /* 默认系统会加 break */
   } 
   ```

2. with no condition

   ```go
   switch {
   case t.Hour() > 10:
   	...
   case t.Hour() > 20:
   	...
   default:
   	...
   }
   ```

### defer
1. execute sequence.FILO
   ```go
func d() {
    defer Expression1;
    defer Expreesion2;
    // stacking defer, 先执行Expression2，再执行Expression1；
}
   ```

2. `defer` in loop

   ```go
   // ❌错误示例
   for {
     f := os.File("...")
     defer f.close() // close会成为最后一个f实例。
   }
   
   // 方案
   for {
     func () {
        f := os.File("...")
     	defer f.close()
     }()
   }
   ```
   
3.  A deferred function's arguments are evaluated when the defer statement is evaluated.
  
  ```go
  func() {
    i := 0
    defer fmt.Println(i) // print: 0
    i++
  }  
  ```
  
4. Deferred functions may read and assign to the returning function's named return values.

   ```go
   func c() (i int) {
       defer func() { i++ }()
       return 1
   }
   // 最终返回的i值为： 2
   ```

   

5. `defer` with `panic`。defer可以再Panic后执行。

   ```go
   defer func() { 
     if r := recover(); r!=nil { // recover 类似于异常被catch，上一级的flow正常执行。
       ...
     }
     panic("raise a Panic")
     ... // 不会被执行到的代码
   }
   // Panic 被Recover后，会继续执行流程。打印: continue...
   ```

   

### slice

1. literals: `s := []int{1, 2, 3, 4}`.
2. *null slice*: length is 0, capability is 0, no underlying array. 

3. slice os slice: `[][]string{[]string{"a", "b"}, []string{"e", "f"}}`
4. *append*: `append(s []T, vs ... T)`

### range

1. to slice. `index, [v] := range []T` *v*可以省略，只返回index ；用`_`代替index或者v
2. to map. `key, [value] := range mapVar`

### map

1. literals: `map[string]int{"key1": "s1", "key2": "s2"}`. `key` is required.

2. literals 2:

   ```go
   map[string]Vertex {
       "vertex1": Vertex{1, 2},
       "vertex2": Vertex{3, 4},
   }
   ```
```

3. `delete(Map, Key)` if Key not in Map, then no op.

4. two-value assignments. `v, ok := Map[Key]` 如果`Key`不存在ok=false,否则ok=true。
```
### Function

1. function as variables.

   ```go
   func varF(x, y int) int {
       return x * y
   }
   func main() {
       f := varF
       f(1, 2)
   }
   ```

   2. closure

   ```go
   func main() {
       x := 1
       func () { // closure function need anonymous function?
           fmt.Println("x", x)
   	}()
   }
   ```

### struct

**Embedding**: If you embed a nameless struct then the embedded struct provides its state (and methods) to the embedding struct directly 

```go
type Person struct {
	name, sex string
}
type Policeman struct {
	Person // Embeding；继承的核心
	station string
}
// 生成实例方法；返回值为*Policeman， 则返回时需要&
func NewPoliceman(name, sex, station string) *Policeman {
	return &Policeman{Person{name, sex}, station} // struct literal
}
func main() {
  /* 指针和非指针最终出来时效果一样。细节指针和非指针区别？？
  */
	police := Policeman{Person{"wang", "male"}, "abc"}
	police2 := &Policeman{Person{"wang", "male"}, "abc"}  
	police3 := NewPoliceman("zhang", "femail", "Peak")

	fmt.Println("Hello, playground", police.name, police2.name, police3.name, police3)
}


```





### Method

***Define:*** a function with a special receiver argument.

***tips:***

1. can only declare a method with a receiver whose type is defined in the same package as the method.

   eg. `type myInt int`才可以定义method。~~(p int)FooFuc(){}~~错误。

2. **value receiver** operate on copy of value.

    > 所以修改receiver的值时只能用指针类型method。

***method and pointer indirection(间接):***

1. pass pointer to pointer method.

   ```go
   type Vertex struct {
     x, y int
   }
   func (v *Vertext) F() {
     ...
   }
   
   func main() {
     v := Vertex{3, 4}
     v.F() // go interpret to (&v).F()
   }
   ```

2. pass value to value method

   ```go
   func (v Vertext) F() {
     ...
   }
   ...
   pv := &Vertex{3, 4}
   pv.F() // go interpret pv to (*pv).F()
   ```

   

### interface

- ***Define:*** a **interface** defined as a set of signatures.

- ***Underlying：*** (value, Type) sets. 

    > 如果赋值给interface 类型I的value没有实现该值的任何方法，则变异抛异常。

- ***nil interface value***: 如果传入的value没有实现interface的method，则传入**nil**指针调用该方法。

    ```go
    type I interface {
      M() [ReturnType]
    }
    ...
    var i I
    i.M()  // panic
    
    ```

- ***interface values with nil underlying values***

  ```go
  type I interface {
    M()
  }
  type MyType int
  func (p *MyType)M() {
    fmt.Println(p)
  }
  ...
  var i I
  var v MyType
  i = v // 区别nil interface value，需要明确i实现的具体类型
  i.M() //  result -> print  "nil"
  ```

- ***Type Assertion***：provide access to an interface value's underlying concret value. 

  >  `v, [ok] = i.(typeNma)` 如果没有OK则抛异常。

- ***Type Switch***

  ```go
  switch v := interfaceValue.(type) {
  case type1:
  		fmt.Println(v)  // v 可以使用。
      ...
  case type2:
  }
  ```

- ***golang广泛（uniquitous）使用的的interface***

  ```go
  // 用于自定义fmt打印，类似python的 `__str__`
  type Stringer interface {
    String() string
  }
  ```

  ```go
  // 实现自定义error的输出
  type error interface {
    Error() string
  }
  ```


### other

- **特殊类型转变**

  如：字符串转字符数组`[]byte(stringValue)`。

- #### tags

  ```go
  type M struct {
    Field int `a: ""` //unmashal field需要captain
  }
  ...
  t = reflect.TypeOf(M)
  f, _ := t.FieldByName("field")
  f.Tag.Lookup(key string) (string, bool)/f.Tag.Get(key string) string() 
  ```


#### Goroutines

- ***gramma:***go f(x, y, z)

- `sync`  access to shared memory synchronized.

#### Channels

- ***gramma:***

  ```go
  ch <- v // Send v to channel ch.
  v := <-ch // Receive from ch, and assign value to v
  ch := make(chan int)  // create channel
  ```

#### select

Blocks util one of its cases can run.use a `default` case to try a send or receive without blocking;

```go
for {
  select {
    case c <- x:
    	...
    case <-quit:
    	return
    default:  // 无defalt为阻塞；有则为非阻塞。
     time.Sleep(500 * time.Millisecond)
  }
}
```

#### Mutex (mutual exclusion) 互斥

- `sync.Mutex.Lock`
- `sync.Mutex.Unlock`



---



