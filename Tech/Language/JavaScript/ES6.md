## ES基本语法

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
> 2. 遍历array。`for of`返回元素，`forin`返回元素索引index。
> 3. `forof`内部通过"iteration"实现。支持Array、Set、Map、Generator、String、ArrayLike（arguments、DomNodeList etc）。

