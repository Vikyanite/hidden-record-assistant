# LOL隐藏战绩查询小助手
## 快速上手

1. 以管理员权限打开IDE
2. 打开并登录LOL客户端

3. 在项目主目录(/hidden-record-assistant)下，终端输入一下代码即可运行项目

``` sh
wails dev
```

## 更新

UPDATE 2023/8.28:

vue小成，总算可以开始写了~

UPDATE 2023/8.21:

不过这个网站时灵时不灵的，于是就狠下心学习了Riot开放的LCUAPI的用法，直接从拳头那里查数据，而不依赖这个网站了，同时也改用Wails框架，这也就是这个分支的目的~

UPDATE 2023/8.22:

大概花了一天时间研究了一下，是否有通过其他接口的方式不用打开LOL客户端，起个后台程序之类的就可以直接使用LCUAPI的方式，很遗憾并没有成功。于是还是同大伙一样，需要登录LOL客户端之后才能运行本程序。不过我还是很好奇，那么WeGame又是怎么做到调用LCUAPI的呢，难道它并没有使用LCUAPI？

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
