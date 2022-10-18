<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [AvatarList 用户头像列表](#avatarlist-%E7%94%A8%E6%88%B7%E5%A4%B4%E5%83%8F%E5%88%97%E8%A1%A8)
  - [代码演示  demo](#%E4%BB%A3%E7%A0%81%E6%BC%94%E7%A4%BA--demo)
  - [API](#api)
    - [AvatarList](#avatarlist)
    - [AvatarList.Item](#avatarlistitem)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# AvatarList 用户头像列表


一组用户头像，常用在项目/团队成员列表。可通过设置 `size` 属性来指定头像大小。



引用方式：

```javascript
import AvatarList from '@/components/AvatarList'
const AvatarListItem = AvatarList.Item

export default {
    components: {
        AvatarList,
        AvatarListItem
    }
}
```



## 代码演示  [demo](https://pro.loacg.com/test/home)

```html
<avatar-list size="mini">
    <avatar-list-item tips="Jake" src="https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png" />
    <avatar-list-item tips="Andy" src="https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
</avatar-list>
```
或
```html
<avatar-list :max-length="3">
    <avatar-list-item tips="Jake" src="https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png" />
    <avatar-list-item tips="Andy" src="https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
    <avatar-list-item tips="Niko" src="https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png" />
</avatar-list>
```



## API

### AvatarList

| 参数               | 说明       | 类型                                 | 默认值       |
| ---------------- | -------- | ---------------------------------- | --------- |
| size             | 头像大小     | `large`、`small` 、`mini`, `default` | `default` |
| maxLength        | 要显示的最大项目 | number                             | -         |
| excessItemsStyle | 多余的项目风格  | CSSProperties                      | -         |

### AvatarList.Item

| 参数   | 说明     | 类型        | 默认值 |
| ---- | ------ | --------- | --- |
| tips | 头像展示文案 | string | -   |
| src  | 头像图片连接 | string    | -   |

