# :bash:

### sed&pear (in-place substitution/原地替换)

1. ``` sed -i "" 's/old_string/new_string/g'  \`grep old_string -rl ./ | grep -vE "tags|svn"\` ```     //特殊要求的替换：此命令中要求过滤掉含有tags和svn的文件。mac 版本-i  *extension* ,以*extension*为后缀保存原文件。
   > 如果文件尾部没有换行符，sed会默认加上，所以该方案不是好的方案。
2. `perl -i [extension] -pe 's/STRING/REPLACEMENT/'`perl版本，避免尾部换行符问题。

### buffer vs cache 定义区别。

2. A `buffer` is something that has yet to be "written" to disk.
   
   > 提前申请的空间，用来写入读取等操作。

3. A `cache` is something that has been "read" from the disk and stored for later use。
   
   > 即缓存。把处理后的数据存储在特定空间（一般性能更好的空间）用来更方便的使用。

### if:

```bash
   # []中为test 命令语法 [ <statement> ]有空格,见*man test*.
   # Boolean Operators || &&
   if [ <some statment> ]  
   then
       <commands>
   elif [<some statment>]
   then
       <else commands>
   else
       <other commands>
   fi
```

### case

```bash
   case <variable> in 
       <pattern1>)
           <commands>
           ;;
   esac
```

### crontab

> crontab网上有`quartz`(java任务框架)，和linux的crontab表达式不同。quartz精确到秒。

#### 格式:

   ​    **minute hour day(of month) month day(of week) COMMAND**

> `day of month`和`day of week`容易冲突；同时选择同时生效。

| field        | allowed value          |
| ------------ | ---------------------- |
| minute       | 0-59                   |
| hour         | 0-23                   |
| day of month | 1-31                   |
| month        | 1-12 or names          |
| day of week  | 0-7 (0 or 7 is Sunday) |

   *(asterisk): 全部，不限；

   m-n(hyphen）: m到n;

   x,y,z: 列表[x, y, z];

   /: 每多久;

   %(percent-sign): COMMAND中如果没有被反斜杠\解释，则%后的内容会标准输入执行。 

### io

> stdin: 0
> stdout: 1
> stderror: 2

   redirection（重定向）` command  [1,2]> FILE(&1,&2)`

```bash
   echo "text" >&FD # 重导向到&FD，>&FD 位置可变。
   echo "text" FD1>&FD2 # 把FD1重定向到FD2.
```

> **FD>FD 不可以有空格。**

### xargs

   `xargs -I REPLACEMENT cmd ...REPLACEMENT` 通过-I指定占位符。

### awk

   `awk -F SEPARATOR '{print $1...$N}'`从`$1`开始。`$0`表示整行。`$NF`最后一行。 

### globbing：文件名拓展 /wildcard :

文件名扩展：先匹配再扩展。如echo ?.txt,如果内容下无匹配，则会输出?.txt,如果有则输出a.txt。对应为通配符/非文件名扩展。 

文件名扩展单引号里时，touch 'foo*',则创建foo*文件

- *,?通配置。（文件名拓展）
- [a-b],[ab],[^a],[!a].(文件名拓展)。touch [x,y].txt如果文件不存在则报no match。
- {x,y,z,zz},先拓展再处理(非文件名拓展)。touch {x,y}.txt 如果文件不存在，则会创建x.txt,y.txt。可嵌套，echo {a..c}.txt=>a.txt,b.txt,c.txt
- 变量拓展：${},${!S}扩展为所有。可嵌套,${X, ${b}}.
- 子命令：$(date).可嵌套（旧的不去写法`date`）。
- 算数扩展$((1+2))。
- 自负类:[[:class:]],**class**如`alpha`，`upper`，`blank`，`digit`,` graph`(a-z,A-Z,1-9),`upper`（A-Z）,`digit`(16进制)等

特殊变量

脚本文件内部，可以使用特殊变量，引用这些参数。

- `$0`：脚本文件名，即script.sh。
- `$1`-`$9`：对应脚本的第一个参数到第九个参数。
- `$#`：参数的总数(数字)。
- `$@`：全部的参数，参数之间使用空格分隔。
- `$*`：全部的参数，参数之间使用变量$IFS值的第一个字符分隔，默认为空格

### Bash

1. bash中空格不能多，如果=号中间没有空格`a=b`为赋值；如果有空格`a = b`,则是条件判断。

2. `set`:Display, set or unset values of shell attributes and positional parameters.

3. `mktemp` 生成临时文件名。

4. `trap`设置处理signal处理。`trap commnds SIGNAL`.

5. `${!VAR}`取变量最终值。如果变量VAR的值也是变量。则获得变量指的变量的值。

6. 字符串操作。
   
   1. `${ #STRING }` 字符串长度。
   
   2. `${ STRING: offset [:length] }` 截取字符串。offset 从0开始。
   
   3. 正则匹配操作。
      
           1. 头部开始匹配删除。`${ STRING#PATTERN }`（非贪婪模式）删除正则匹配内容。`${ STRING##PATTERN}`贪婪模式。*如果匹配不成功则返回原字符串。*
      
      2. 尾部开始匹配删除。`${ STRING%PATTERN }`非贪婪模式。`${ STRING%%pattern}`贪婪模式。
         
               3. 任意位置开始。`${ STRING/PATTERN}`正则模式。`${ STRING//PATTERN}`f
               4. 替换。`${ STRING/[#/%]PATTERN/SUBSTITUTION}`
   
   4. 改变大小写。 `${ STRING^^ }`大写。`{ STRING,, }`小写。zsh下`${ STRING:u}` uppercase和`${ STRING:l}` lowercase。

7. `read [-s(密码模式)] [-p "prompt"] [-d TERMINATOR(单字符，缺省回车键)] [-n NUM] [VAR]`输入内容给变量VAR，如果VAR未提供，则赋值到环境变量`REPPLY`.

8. 数组。`a=(1 2 3)` 。`a[INDEX]`index从1开始.     

### find

`find -type [f/d]` 指定文件类型 f(ile) or d(irectory)

#### `shebang`:

*If you're writing scripts, ensure it has a [shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)) at the top*

#### glob vs regexp

*regexp 比glob语法复杂，实现更多的功能。*

1. glob：实现的语法比较少。主要：`[]`,`[^]`,`?`,`*`。

2. 含义不太一样，对应关系：`? - .; * - .*`

3. 核心差异没有实现`克林闭包`。
   
   > 克林闭包（Kleene star or Kleene operator or Kleene closure）：R*语法。
   > 
   > 即a\*    表示:a, aa, aaa

-----

1. 打印文件第m-n行。
   
   `awk 'FNR>=20 && FNR<=40' file_name`I
