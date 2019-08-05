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
> 	1. 正常`raise Exception`;
>
>  	2. 直接` raise`，reraise the *current exception*in an exception handler ;

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

`yield`: yield



### Derector

