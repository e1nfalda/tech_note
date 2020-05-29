### JSX
```javascript
1. `<TAG>...*</TAG>`
```

### custom components
```js
import React, { Component } from 'react'
// class
export  [default] class COMPONENT_NAME extends Component {
  render() {
    return (JSX)
  }
}
// function
export [default] function COMPONENT_NAME () {
  return ( JSX )
}
```

### props：
```js
function Cat(props) {
  return (<View>{ props.name }</view>)
}

<CAT name="abc">
```

### state
```js
// function
[ VAR_NAME, SETTER ] = setState(INIT_VALUE)  // 声明；setState: 钩子方法
...
<tag onClick={ SETTER(new_value) } >{ VAR_NAME }</tag>

// class

state = { VAR_NAME: INIT_VALUE}
...
<tag onClick=setState({VAR_NAME: NEW_VALUE}) >{ this.state.VAR_NAME }</tag>

```

### 不同平台代码不同时

#### Platform-specific module

```js 
import { Platform } from "react-native"
Platform.OS  // 'ios' | 'android' 
...Platform.select({
	"ios": {...},
    "android": {...},
    ...
}) // 'ios' | 'android' | 'native' | 'default'
```

#### Platform-specific extensions

文件名不同。

`MODULE_NAME.[ios/android].js`,使用时直接引入`import ... from 'MOUDULE_NAME'`。

