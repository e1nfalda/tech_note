## bash

1. ``` sed -i 's/old_string/new_string/g'  \`grep old_string -rl ./ | grep -vE "tags|svn"\` ```     //特殊要求的替换：此命令中要求过滤掉含有tags和svn的文件。

2. A `buffer` is something that has yet to be "written" to disk.
   
> ？一个空间可以把食物累加到一起再批量操作。

3. A `cache` is something that has been "read" from the disk and stored for later use。

   > ？为了更快速的被使用，把内容存储在更快速的层级。

4. ### if:

   ```bash
   # []中为test 命令语法 [ <statement> ]有空格。见man test；Boolean Operators || &&
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

5. ### case

   ```bash
   case <variable> in 
   	<pattern1>)
   		<commands>
   		;;
   esac
   ```

6. ### crontab

   > crontab网上有`quartz`(java任务框架)，和linux的crontab表达式不同。quartz精确到秒。

   #### 格式:

   ​	**minute hour day(of month) month day(of week) COMMAND**

   > `day of month`和`day of week`容易冲突；同时选择同时生效。

   | field  | allowed value |
   | ------ | ------------- |
   | minute | 0-59          |
   |hour|0-23|
   |day of month|1-31|
   |month|1-12 or names|
   |day of week|0-7 (0 or 7 is Sunday)|

   *(asterisk): 全部，不限；

   m-n(hyphen）: m到n;

   x,y,z: 列表[x, y, z];

   /: 每多久;

   %(percent-sign): COMMAND中如果没有被反斜杠\解释，则%后的内容会最会标准输入执行。 
   
7. ### stdio

   >stdin: 0
   >
   >stdout: 1
   >
   >stderror: 2

   redirection（重定向）

   ```	bash
   echo "text" >&FD # 重导向到&FD，>&FD 位置可变。
   echo "text" FD1>&FD2 # 把FD1重定向到FD2
   ```

