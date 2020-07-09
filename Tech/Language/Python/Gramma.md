## 语法

### with

​	 [pep343](https://www.python.org/dev/peps/pep-0343/)

​	`__exit__(self, exc_type, exc_value, traceback)`: 注意参数！如果方法返回*True*,异常不会抛出。

```python
class WithDemo(object):
    def __enter__(self):
        print "enter"
        return self
    def __exit__(self, exc_type, exc_value, traceback):
        print "excption", str(exc_type), str(exc_value), str(traceback)
        return bool_value // bool_value为False时会把异常重新抛出，
    def pf(self):
        print "With Demo"
    def raise_exc(self):
        raise Exception("Exception")
        
with WithDemo() as f:
    f.pf()
    f.raise_exc()
```

### raise

> raise有两种场景。
>
> 1. 正常 `raise Exception`;
>
> 2. 直接` raise`无参数时，reraise the *current exception*in an exception handler ;可参照`with`语法。
>   ```python
>      try:
>      	raise MyException("efg")
>      except Exception as e:
>         raise  # MyException会被再次抛出
>   ```
>
> ​     
### Exception

作用：

1. 父级Exception可以`except`其子类Exception。
2. 多个Exception如果都可以匹配，匹配到先碰到的。

```python
class Level1Exception(Exception):
  def __init__(self, value):
    self.value = value
  def __str__(self):
    return "level1 error: {}".format(self.value) 
  
class Level2Exception(Level1Exception):
  pass

try:
  raise Level2Exception("level2 Exception")
except Level1Exception:  # Level1Exception,Level2Exception 先碰到哪个则执行哪个
  print "level1"
except Level2Exception:
  print "level2"
```
### Iterator(迭代器)

`iter()`：把一个可以迭代的对象**变为**迭代器对象。对象有`__iter__()`方法。

`StopIteration`: 终止迭代抛出的异常。

`__next__`: 迭代调用的方法。

`next()`: 执行一次迭代。

```python
class RevStr:
  """ 字符串反转
  """
  def __init__(string):
    self.s = string
    self.index = len(s)
  def __iter__(self):
    return self
  def __next__(self):
    self.index -= 1
    if self.index < 0:
      raise StopIteration
 		else:
      return self.s[self.index]
```



- iterator Class不能是newstyle class（继承object）。

### Generator(生成器)

`yield`: 使用yield的场景即Generator。

### Decorator(装饰器)

> 装饰器调用关系：decorate().fn()

```python
def derector(fn):
  def wrapper(*args, **kwargs):
    ...fn(*args, **kwargs)..
  return wrapper

# 带参数
def decorator(dec_arg...):
  def wrapper(fn):
    @functoos.wraps(fn)  # 包到最终返回的函数上
    def f(*args, **kwargs)
      ...
      fn(*args, **kwargs)
      ...
    return f
  return wrapper

# class
class Decorator():
  def __init__(self, fn):
    functools.update_wrapper(self, fn)
    self.fn = fn
  def __call__(self, *args, **kwargs):
    self.fn(*args, **kwargs)
```

### type()

1. type(object) 返回object的type。

   ```python
   class MyClass(object):
     pass
   mc = MyClass()
   type(mc) == MyClass
   ```

   

2. type(name, bases, dict) 生成一个新type。

   ```python
   X = type('X', (object, ), dict(a=1)) 
   # 等效:
   class X(object):
   #  a = 1
   ```

### method调用

1. 实例直接调用。`instance.fn(*args)`

2.  通过对象调用。`Object.fn(instance, *args)。`

   > *with*语法实现原理中`enter = (type(object).__enter__)`
   
### file
#### mode

*read模式不能write，write模式不能read。*

- r：缺省值。如果文件不存在IOError;
- w: 如果存在则删除文件后再新建。
- a: append.文件存在打开，不存在新建。
- b: 不做特殊字符转换。如果没有比如换行符会换成平台对应的换行符。默认为t.
- w+,a+,r+: updating。w+也会先删除。

#### performance

`read([size])`: 读取size或全部内容为string。大文件可能内存溢出等。

`readlines()`: 读取全部文件为string list。

`readline():`iterator, 读取一行。

`for line in fd`: 系统缓存优化，速度快。==推荐方式==



### 其他：

**inner class:** class 中定义class,为了更好的代码结构。通过`six`package with_classmeta。
   ```


   ```