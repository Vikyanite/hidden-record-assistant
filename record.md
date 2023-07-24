## 开发过程记录

7.21

我们通过F12去查看，我们到底发送了什么包

![image-20230721185848408](C:\Users\vic\AppData\Roaming\Typora\typora-user-images\image-20230721185848408.png)

结果神奇的发现，这网站居然是用PHP写的。不过没关系，我们忽略掉图片之后，很容易就能找到这个网站发出的请求是什么，也就是一个拥有`{查询名字，大区编号，起始搜索位置，结束搜索位置，搜索类型（LOL/云顶），SC}`这六个字段的表单。

其中`SC`字段让我很不理解是什么，不过不着急，我们先用APIPost模拟一下请求。继续在浏览器中查看刚刚请求发送的方式，发现是`POST`

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-b1bcb2-image-20230721190151428.png)

然后我们就打开APIPost输入对应的字段:

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-ea9479-image-20230721190408426.png)

然后发现出错了：

![image-20230721190804931](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-00efac-image-20230721190804931.png)

看来`SC`字段是必传的，有可能是token那一类的，于是就去查PHP Token，但是没啥结果，

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-f7b243-image-20230721184544641.png)

