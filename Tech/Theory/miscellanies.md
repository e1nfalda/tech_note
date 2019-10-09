## miscellanies

### EBNF（扩展的巴科斯范式）

> *metasyntax*一种。和`BNF`语法稍有差异。python使用的BNF,golang使用的EBNF。[wike]([https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form))

|                            Usage                             | Notation  |
| :----------------------------------------------------------: | :-------: |
|                          definition                          |     =     |
| [concatenation](https://en.wikipedia.org/wiki/Concatenation) |     ,     |
|                         termination                          |     ;     |
| [alternation](https://en.wikipedia.org/wiki/Alternation_(formal_language_theory)) |    \|     |
|                           optional                           |  [ ... ]  |
|                          repetition                          |  { ... }  |
|                           grouping                           |  ( ... )  |
|                       terminal string                        |  " ... "  |
|                       terminal string                        |  ' ... '  |
|                           comment                            | (* ... *) |
|                       special sequence                       |  ? ... ?  |
|                          exception                           |     -     |

### 鸭子类型（duck typing）

[wiki]([https://zh.wikipedia.org/wiki/%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B](https://zh.wikipedia.org/wiki/鸭子类型))

#### 定义: 

  程序设计中动态类型的一种风格。对象的有效语义不是继承的类或者实现特定接口，而是由“**方法**和**属性**的集合”决定。应用的理解中，可以不使用继承的情况下实现多态。

#### 该风格的语言：

  python的类调用

  ```python
"""
因为字符或者int都具有相应的__add__方法，调用时只要是实现了该方法就可以动态调用。或者如python 的file()、cString 模块
"""
def sum(x, y): 
  return x+y

sum(1, 2) 
sum("x", "y")

  ```

golang，只要实现了相应接口的方法就可以调用。

#### 其他风格：

**接口：**需要实现接口中下定义的所以方法。 



