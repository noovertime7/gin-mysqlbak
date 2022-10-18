<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [IconSelector](#iconselector)
    - [使用方式](#%E4%BD%BF%E7%94%A8%E6%96%B9%E5%BC%8F)
    - [事件](#%E4%BA%8B%E4%BB%B6)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

IconSelector
====

> 图标选择组件，常用于为某一个数据设定一个图标时使用
> eg: 设定菜单列表时，为每个菜单设定一个图标

该组件由 [@Saraka](https://github.com/saraka-tsukai) 封装



### 使用方式

```vue
<template>
	<div>
       <icon-selector @change="handleIconChange"/>
    </div>
</template>

<script>
import IconSelector from '@/components/IconSelector'

export default {
  name: 'YourView',
  components: {
    IconSelector
  },
  data () {
    return {
    }
  },
  methods: {
    handleIconChange (icon) {
      console.log('change Icon', icon)
    }
  }
}
</script>
```



### 事件


| 名称   | 说明                       | 类型   | 默认值 |
| ------ | -------------------------- | ------ | ------ |
| change | 当改变了 `icon` 选中项触发 | String | -      |
