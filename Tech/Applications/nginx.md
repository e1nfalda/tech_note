### location

语法：**location [= | ~ | ~\* | ^~] uri {...}**   / **location @name {...}**

> prefix-based location/label: 没有modifier,=,~^。 
>
> regex-based location: 正则类型。~*, ~



**modifier**:

1. **=** ：精准匹配。
2. **～**：正则匹配。区分大小写。
3. **～***：正则匹配。不区分大小写。
4. **^~**: 以该标签开头的字段。

**匹配流程**：

- 先匹配*prefix-based location*。

  1. 如果=号并且完全匹配则停止匹配。
  2. 如果匹配到^~的规则，则选择最长的匹配（longest match），并终止匹配。
  3. 如果缺省modifier匹配则记录下匹配的最长的（longest match）字段。

- 匹配正则规则。

  2. 如果正则匹配到则使用匹配到的最早（first match）的规则。

  3. 如果未匹配到而有prefix匹配到则从prefix中选最长的匹配规则。

     