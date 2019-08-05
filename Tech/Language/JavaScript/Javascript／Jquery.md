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


