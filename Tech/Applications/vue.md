# vue

## 响应式原理

### 变化追踪

Object.defineProperty把所有属性转为getter/setter

### 流程图

### Object类型

通过特定方法设置
Vue.set/this.$set(object, key, value)
 重新生成对象。
this.someObject = Object.assign({}, this.someObject)

### array
Vue.set(array, index, newValue)
array.splice(index, 1, newValue)
 **支持的方法**: `push()`, `pop()`,`shift()`, `unshift()`, `splice()`, `sort()`, `reverse()`

### 声明响应式属性。

Vue 不允许动态添加根属性，所以属性需要初始化时声明。

### 异步更新队列

vm.$nextTick(callback)

## js

### [EventLoop](http://www.ruanyifeng.com/blog/2014/10/event-loop.html)

- setTimeout
- setInverval
- process.nextTick(Node.js)
- setImmediate(Node.js)

## re-render component

### reload-page

### v-if

### force-update

### :key

