# LOL隐藏战绩查询小助手
## 快速上手

1. 以管理员权限打开IDE
2. 打开并登录LOL客户端

3. 在项目主目录(/hidden-record-assistant)下，终端输入一下代码即可运行项目

``` sh
wails dev
```

## TODO

- 查询结果界面将英雄等级作为小圆标显示在英雄头像旁边
- 将每个战绩添加牛马值显示并根据牛马值显示对应title
- 完善统计数据（位置选用率、英雄选用率、最近牛马值）
- 多人战绩查询
- 开黑特殊标记
- 主题优化（日间/夜间皮肤切换...
- bug与case修复
- 查询效率提高

## 已实现

- 战绩查询
- 详细对局显示
- 部分最近数据统计
- 基础样式敲定

## 前言

之前因为听说查战绩能提高单排胜率，于是就百度找到了隐藏战绩查询的网页（非常好用~）

但是一直以来有几个痛点就是：

- 每次只能一个人一个人查
- 而且网站失灵时不灵的，最好是能多个网站一起查询
- 希望能过滤掉不是排位的战绩

于是这款帮助我上分的小助手就顺应而出啦~

## 项目运行要求环境：

- GCC
- Go >= 1.13
- wails v2

## 主要功能

其实就是根据上面的痛点推出的功能啦~

- 可以一次性查询房间内的**己方**所有人
- 可以过滤掉不是排位的战绩
- 可以多个网站一起查询，从而保证能查到这个人的战绩

## 特别感谢



## 参考项目/教程：

[real-web-world/hh-lol-prophet: lol 对局先知 上等马 牛马分析程序](https://github.com/real-web-world/hh-lol-prophet)

[上等马还是下等马，英雄联盟LCUAPI研究](https://cloud.tencent.com/developer/article/1987709)

界面与API参考：[Zzaphkiel/Seraphine: 基于 LCU API 的英雄联盟战绩查询工具](https://github.com/Zzaphkiel/Seraphine)

界面参考：[LOLstatsVUE: Full league of legends props.Datas app (like OP.GG)](https://github.com/tudorcrisan1231/LOLstatsVUE/tree/master)
