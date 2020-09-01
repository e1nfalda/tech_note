### 修饰符

| 修饰符      | 当前类 | 同一包内 | 子孙类(同一包) | 子孙类(不同包)   | 其他包 |
| :---------- | :----- | :------- | :------------- | :--------------- | :----- |
| `public`    | Y      | Y        | Y              | Y                | Y      |
| `protected` | Y      | Y        | Y              | Y/N | N      |
| `default`   | Y      | Y        | Y              | N                | N      |
| `private`   | Y      | N        | N              | N                | N      |

#### 访问控制

> - 父类中声明为 public 的方法在子类中也必须为 public。
> - 父类中声明为 protected 的方法在子类中要么声明为 protected，要么声明为 public，不能声明为 private。
> - 父类中声明为 private 的方法，不能够被继承。

default,public,private， protected。

##### protected:

1. 同一包时子类可访问所有。

2. 不同时：

   [^protected]: 子类实例可以访问其从基类继承而来的类方法，而不能访问基类实例的protected方法。

#### 非访问控制

final, static, abstract,synchronous,volatile,transient

### 变量

|        |    局部变量    |                        实例变量                         |   类变量   |
| :----: | :------------: | :-----------------------------------------------------: | :--------: |
| 初始化 | 必须手动初始化 | 默认初始化 <br>（数值类型：0，bool： false, ref: null） | 同实例变量 |
|        |       栈       |                                                         |   静态区   |



#### 局部变量

1. 局部变量没有默认值.	

#### 类变量（静态变量）

#### 成员变量（非静态变量）

### 枚举类型

`enum variableName {enum1, enum2...}`

### 源文件声明规则

在本节的最后部分，我们将学习源文件的声明规则。当在一个源文件中定义多个类，并且还有import语句和package语句时，要特别注意这些规则。

  * 一个源文件中只能有一个 public 类
  * 一个源文件可以有多个非 public 类
  * 源文件的名称应该和 public 类的类名保持一致。例如：源文件中 public 类的类名是 Employee，那么源文件应该命名为Employee.java。
  * 如果一个类定义在某个包中，那么 package 语句应该在源文件的首行。
  * 如果源文件包含 import 语句，那么应该放在 package 语句和类定义之间。如果没有 package 语句，那么 import 语句应该在源文件中最前面。
  * import 语句和 package 语句对源文件中定义的所有类都有效。在同一源文件中，不能给不同的类不同的包声明。

类有若干种访问级别，并且类也分不同的类型：抽象类和 final 类等。这些将在访问控制章节介绍。

### 类型


#### 内置

- **byte**: 8bit.

- **short**: 16 bit.

- **int**: 32bit

- **long**: 64bit. *12L*

- **float**: 单精度浮点。32bit。`1.23f`

- **double**: 双精度浮点。64bit。`1.23d`

- **boolean**: *default: `false`*

- **char**: *unicode*.16bit.

#### 引用变量

> - Array、Object等。
>
> - default null。

### 常量

`final TYPE var = DATA`

## 结构

### 循环

#### while

```java
while (BOOLEAN EXPRESSION) {
    ...
}
```

#### do...while

```java
do {
    ...
} while (BOOLEAN EXPRESSION)
```

#### for

```java
for (INIT, CODITION, UPDATE) {
    ...
}
```

#### for(enhance)

```java
for (TYEP VAR : ARRAY|FUNC()[TYPE]) { // FUNC: 返回值为数组
}
```

#### BREAK;CONTINUE;

### 条件：

#### if

```java
if (BOOLEAN_EXPRESS) {...} 
[else if (BOOLEAN_EXPRESSION) {...}]
[else {}]
```



#### switch case

```java
switch (EXPRESSION) {  // EXPRESS: 数字类型/char。SE>7支持String。
    case NUM|STRING:  // 数字类型或字符串常量或字面量。
        ...
        break;
    default:
        ...
}
```

### 继承

**单继承**

`extends`、`implements`

`super`

`this`

`final`

#### 构造方法

- 方法名同类名的方法。
- 会默认调用无参数的父类。

### 范型（generics）

#### 范型方法

#### 范型类

### package

`import package[.package...].[className|*]`

`package com.package.fileName` => path:*com/package/fileName.java*




-----

### other:

1. java 数字类型长度与系统无关。int不论32位还是64位系统，都是4个字节。
2. 整型、实型（常量）、字符型数据可以混合运算。运算中，不同类型的数据先转化为同一类型，然后进行运算。高类型转底是需要强转`(TO_TYPE)var`（内容可能溢出丢失）。

```java
// 低  ------------------------------------>  高
byte,short,char—> int —> long—> float —> double 
```

3. [内存结构](https://deepu.tech/memory-management-in-jvm/)。

   ![JVM Memory structure](./.src/8uh8SPy.png)



[https://baidu.com]: 
