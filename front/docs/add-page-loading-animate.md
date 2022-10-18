<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [为首屏增加 加载动画](#%E4%B8%BA%E9%A6%96%E5%B1%8F%E5%A2%9E%E5%8A%A0-%E5%8A%A0%E8%BD%BD%E5%8A%A8%E7%94%BB)
  - [需求](#%E9%9C%80%E6%B1%82)
  - [实现方案](#%E5%AE%9E%E7%8E%B0%E6%96%B9%E6%A1%88)
  - [写在最后](#%E5%86%99%E5%9C%A8%E6%9C%80%E5%90%8E)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

为首屏增加 加载动画
====



## 需求

> 为了缓解用户第一次访问时，加载 JS 过大所导致用户等待白屏时间过长导致的用户体验不好，进行的一个优化动效。



## 实现方案

1. 将 动画加载 dom 元素放在 #app 内，Vue 生命周期开始时，会自动清掉 #app 下的所有元素。
2. 将 动画加载 dom 元素放在 body 下，Vue 生命周期开始时 App.vue (created, mounted) 调用 `@/utils/utll` 下的 removeLoadingAnimate(#id, timeout) 则会移除加载动画

最后一步：
​	将样式插入到 `public/index.html` 文件的 `<head></head>` 最好写成内联 `<style>动画样式</style>` 



----

目前提供有两个样式，均在 `public/loading` 文件夹内。且 pro 已经默认使用了一套 loading 动画方案，可以直接参考 `public/index.html`


## 写在最后

目前 pro 有页面 overflow 显示出浏览器滚动条时，页面会抖动一下的问题。

欢迎各位提供能解决的方案和实现 demo。如果在条件允许的情况下，建议请直接使用 pro 进行改造，也欢迎直接 PR 到 pro 的仓库
