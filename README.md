## note

### moto 

1. 元认知（metacongnition）
2. 好奇心
3. 简化问题。

4. 思维跳跃性

5. 对观点的态度
6. 对人的态度

### Fragment

### Data Retrieval

**Data retrieval** means obtaining data from a database management system such as ODBMS. In this case, it is considered that data is represented in a structured way, and there is no ambiguity in data.

#### semaphore（信号量） vs mutex

互斥锁和**binary semaphore**比较类似，但机制有不同。互斥锁可以被其他unlock。

#### Bash

1. bash中空格不能多，如果=号中间没有空格`a=b`为赋值；如果有空格`a = b`,则是条件判断。

2. `set`:Display, set or unset values of shell attributes and positional parameters.

3. `mktemp` 生成临时文件名。

4. `trap`设置处理signal处理。`trap commnds SIGNAL`.

5. `${!VAR}`如果变量VAR的值也是变量。则获得变量指的变量的值。

6. 字符串操作。

    1. `${ #STRING }` 字符串长度。

    2. `${ STRING: offset [:length] }` 截取字符串。offset 从0开始。

    3. 正则匹配操作。

        	1. 头部开始匹配删除。`${ STRING#PATTERN }`（非贪婪模式）删除正则匹配内容。`${ STRING##PATTERN}`贪婪模式。*如果匹配不成功则返回原字符串。*
     	2. 尾部开始匹配删除。`${ STRING%PATTERN }`非贪婪模式。`${ STRING%%pattern}`贪婪模式。
              	3. 任意位置开始。`${ STRING/PATTERN}`正则模式。`${ STRING//PATTERN}`f
              	4. 替换。`${ STRING/[#/%]PATTERN/SUBSTITUTION}`
    4. 改变大小写。 `${ STRING^^ }`大写。`{ STRING,, }`小写。zsh下`${ STRING:u}` uppercase和`${ STRING:l}` lowercase。
    
7. 数据重定向。 command  [1,2]> FILE(&1,&2)

8. `read [-s(密码模式)] [-p "prompt"] [-d TERMINATOR(单字符，缺省回车键)] [-n NUM] [VAR]`输入内容给变量VAR，如果VAR未提供，则赋值到环境变量`REPPLY`.
   
9. 数组。`a=(1 2 3)` 。`a[INDEX]`index从1开始。

    ​     

    ​      

### Four Quadrant（四象限）

| 重要性/紧急性 | 不紧急 | 不重要 |
| ------------- | ------ | ------ |
| 不重要        |        |        |
| 重要          |        |        |
