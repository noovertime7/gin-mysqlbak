<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Trend 趋势标记](#trend-%E8%B6%8B%E5%8A%BF%E6%A0%87%E8%AE%B0)
  - [代码演示  demo](#%E4%BB%A3%E7%A0%81%E6%BC%94%E7%A4%BA--demo)
  - [API](#api)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Trend 趋势标记

趋势符号，标记上升和下降趋势。通常用绿色代表“好”，红色代表“不好”，股票涨跌场景除外。



引用方式：

```javascript
import Trend from '@/components/Trend'

export default {
    components: {
        Trend
    }
}
```



## 代码演示  [demo](https://pro.loacg.com/test/home)

```html
<trend flag="up">5%</trend>
```
或
```html
<trend flag="up">
    <span slot="term">工资</span>
    5%
</trend>
```
或
```html
<trend flag="up" term="工资">5%</trend>
```


## API

| 参数      | 说明                                      | 类型         | 默认值 |
|----------|------------------------------------------|-------------|-------|
| flag | 上升下降标识：`up|down` | string | - |
| reverseColor | 颜色反转 | Boolean | false |

