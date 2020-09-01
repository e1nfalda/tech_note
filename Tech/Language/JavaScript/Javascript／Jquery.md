## Javascript／Jquery

### 其他：／

1. dom to jquery object。 j_ele = $(dom)[0]

2. div.natureHeight/natureWidth;

3. `document.ready` is a jQuery event, it runs when the DOM is ready, e.g. all **elements** are there to be found/used, but not necessarily all *content*.

   `window.onload` fires later (or at the same time in the worst/failing cases) when **images** and such are loaded, so if you're using image dimensions for example, you often want to use this instead.

4. html 渲染流程。http://taligarsiel.com/Projects/howbrowserswork1.htm#Render_tree_construction

5. $(html).load 加载页面

6. js复制.
   ```javascript
   $("body").append($temp);
   $temp.val($("#invite_code").text()).select();
   document.execCommand("copy");
   $temp.remove();
   ```

7. `import {X as AliasX} from 'package'`
8. js Boolean.
   - **False**: "", 0, undefined, Nan, null。 
   - **True**: [], {}, Object。
9. 正则。
   - (abc|d) 匹配 “abc”或者“d”。
   - [abc|d] 匹配 "a", "b", **"|"**, "d"。
10. According to the [ES5 spec](http://es5.github.com/), when doing [bitwise operations](http://es5.github.com/#x11.10), all operands are converted to [`ToInt32`](http://es5.github.com/#x9.5) (which first calls [`ToNumber`](http://es5.github.com/#x9.3). If the value is `NaN` or `Infinity`, it's converted to `0`).
