## ES基本语法

### 数据类型

原始数据类型：undifined， null，Boolean，String，Number，Object，Symbol（ES6）

### Destructing（解构）

#### 数组解构造

1. `[x, y, ...z] = [1, 2, 3, 4, 5]` z = [3, 4, 5]。
2. `[x, y, ..., z] = [1, 2, 3, 4, 5]` z = 5。
3. `[x, [y], z] = [1, [2], [3]]` y=2, z=[3]。

#### Object

1. `let value;{"key": value} = {"key": 123}`value=123。
2. `let foo,bar; {foo, bar} = {'foo': 1, 'bar': 2}` foo=1, bar=2。
3. `let value; {"key": value="default"}={}` value="default"。

### 箭头函数（）=> Express

> 箭头函数中this的作用域同*父作用域*。

### 生成器Generator

### for语句

> `for...of`和`for...in`区别：
>
> 1. 遍历object。`for of`无法遍历；`for...in`返回Object key。
> 2. 遍历array。`for of`返回元素，`for...in`返回元素索引index(String)。
> 3. `for...of`内部通过"iteration"实现。支持Array、Set、Map、Generator、String、ArrayLike（arguments、DomNodeList etc）。

### Symbol

> 标示*独一无二的值*。[链接]([http://es6.ruanyifeng.com/#docs/symbol#Symbol-for%EF%BC%8CSymbol-keyFor](http://es6.ruanyifeng.com/#docs/symbol#Symbol-for，Symbol-keyFor))
>
> 1. 可以避免字符串命名的属性可能会冲突。
>
> 2. 声明：`Symbol([description string])`。注：不可以用*new*。
>
> 3. *for...in*、*for...of*，*Object.keys()*,*Object.getOwnPropertyNames*,*JSON.stringify*都不会返回，可以通过*Object.getOwnPropertySymbols()*返回，不属于私有属性。
>
> 4. `Symbol.for(symbolName)`: 定义symbol。区别Symbol()，*for*方法注册全局可以查询。
>
>    `symbolName symbol.keyFor(symbolVar）`：返回Symbol变量用for声明的*symbolName*。
>
> 5. javascript built-in symbols：Symbol. 'iterator| match|hasInstance, ()等.见[devdocs](https://devdocs.io/javascript/global_objects/symbol)**Well-known symbols**。

```javascript
let mySymbol = Symbol()
let a = {};
// 写法。
a[mySymbol] = "hello"  // 第一种写法
a = {[mySymbol]: "hello"} // 第二张写法
Object.defineProperty(a, mySymbol, {value: "hello"})  // 第三种写法
```

### Map

> ES6新增。与Object区别，Object `Key`只能字符串类型（各种对象强转为string），Map可以任意类型。
