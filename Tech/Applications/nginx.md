### location

语法：**location [= | ~ | ~\* | ^~] uri {...}**   / **location @name {...}**

> label-based location: 没有modifier或者=号 
>
> regex-based location: 正则类型。~*, ^~



**modifier**:

1. **=** ：精准匹配。
2. **～**：正则匹配。区分大小写。
3. **～**：正则匹配。不区分大小写。
4. **^~**: 以该uri开头的。

**匹配流程**：

- 先匹配*label-based location*。

  1. 如果=号并且完全匹配则停止匹配。
  2. 如果缺省modifier匹配则记录下来。

- 匹配正则规则。

  1. 如果是以^~开始的规则，匹配到则终止匹配。

  2. 如果正则匹配到则使用匹配的最长的规则。

  3. 如果未匹配到而有label匹配到则从label中选最长的匹配规则。

     