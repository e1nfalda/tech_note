### 面向对象特点:

> 1: 抽象；2: 封装；3: 继承； 4: 多态.
>
> emthod，attribute对象 

### 原型链（prototype chain）

> Each object has a private property which holds a link to another object called its **prototype**. That prototype object has a prototype of its own, and so on until an object is reached with `null` as its prototype. By definition, `null` has no prototype, and acts as the final link in this **prototype chain**.
```javascript
function f() {
  this.a = 1
  this.b = 2
}
let o = new f();
f.prototype.b = 3 
f.prototype.c = 4
// 不可以用 f.prototype = {b: 3, c: 4};会打破原型链
// 原型链： {a: 1, b: 2} --> {b: 3, c: 4} --> Object.prototye --> null
o.a // 1
o.b // 2  b=3属性为"属性遮蔽（property shadowing）"
o.c // 4  如果原型上没有找到该属性，则去上册查找。
o.d // undefined
```
#### 定义
- `__proto__`是浏览器对标准`getPrototypeO`和`setPrototypeOf`的访问对象原型(prototype)实现；访问的是

- *函数(function)*又一个特殊的属性`prototype`。
   > [`Function`](https://devdocs.io/javascript/global_objects/function) objects inherit from `Function.prototype`. `Function.prototype` cannot be modified.
  ```javascript
  function f() {return 2} // 原型链：f --> Function.prototype --> Object.prototype --> null
  ```

 


### 对象创建

#### 最直接方法
> 代码无复用

```javascript
obj1 = new Object();
obj1.name = "name1"
obj1.showName = function() {
  console.log("name:" this.name)
}
```

#### 工厂模式

> 减少了代码创建重复量，但方法对象没有复用。

```javascript
function createPerson(name) {
  var obj  new Object()
  obj.name = name;
  obj.showName = function() {
    console.log("name: ", name)
  }
  return obj
}
person1 = CreatePerson("name1")
person2 = CreatePerson("name2")
person1.showName == person.showName // false
```

#### 构造函数模式

> 同工厂方法，方法对象没有符用

```javascript
function CreatePerson(name) {
  this.name = name
  this.showName = function() {
    console.log("name: ", this.name)
  }
}
person = new CreatePerson("Einfalda")
```

#### 原型（prototype）

> 解决对象不可复用问题

```javascript
function Person(name){
  this.name = name
}
CreatePerson.prototype.ShowName = function() {
  console.log("name:", this.name)
}
```

#### ES6 Class

> es6 支持的语法。

```javascript
class Person {
  // new 时调用的方法 constructor 方法
  constructor(name) { 
    this.name = name
  }
  showName() {
    console.log("name:", this.name)
  }
  // 
  set name(value) {
    console.log("set name")
    this.name = value
  }
  get name() {
    console.log("get name")
    return this.name
  }
}
```

：[msdn](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Inheritance_and_the_prototype_chain)