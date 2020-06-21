#### open: mode

- `r`：缺省值。如果文件不存在`IOError`;

- `w`: 如果存在则删除文件后再新建。

- `a`: append.文件存在打开，不存在新建。

- `b`: 不做特殊字符转换。如果没有比如换行符会换成平台对应的换行符。默认为`t`.

- `w+`,`a+`,`r+`: updating。w+也会先删除

  read模式不能write，write模式不能read。



#### `shebang`: If you're writing scripts, ensure it has a [shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)) at the top

#### glob vs regexp

regexp 比glob语法复杂，实现更多的功能。

1. glob：实现的语法比较少。主要：`[]`,`[^]`,`?`,`*`。

2. 含义不太一样，对应关系：**?**=> **.**; *=>".*"

3. 核心差异没有实现`克林闭包`。

   >
   >
   >克林闭包（Kleene star or Kleene operator or Kleene closure）：R*语法。
   >
   >即a\*	表示:a, aa, aaa