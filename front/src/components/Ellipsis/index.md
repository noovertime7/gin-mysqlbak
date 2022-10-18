<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Ellipsis 文本自动省略号](#ellipsis-%E6%96%87%E6%9C%AC%E8%87%AA%E5%8A%A8%E7%9C%81%E7%95%A5%E5%8F%B7)
  - [代码演示  demo](#%E4%BB%A3%E7%A0%81%E6%BC%94%E7%A4%BA--demo)
  - [API](#api)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Ellipsis 文本自动省略号

文本过长自动处理省略号，支持按照文本长度和最大行数两种方式截取。



引用方式：

```javascript
import Ellipsis from '@/components/Ellipsis'

export default {
    components: {
        Ellipsis
    }
}
```



## 代码演示  [demo](https://pro.loacg.com/test/home)

```html
<ellipsis :length="100" tooltip>
        There were injuries alleged in three cases in 2015, and a
        fourth incident in September, according to the safety recall report. After meeting with US regulators in October, the firm decided to issue a voluntary recall.
</ellipsis>
```



## API


参数 | 说明 | 类型 | 默认值
----|------|-----|------
tooltip | 移动到文本展示完整内容的提示 | boolean | -
length | 在按照长度截取下的文本最大字符数，超过则截取省略 | number | -