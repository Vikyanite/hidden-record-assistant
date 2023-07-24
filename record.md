## 开发过程记录

### 7.21

我们通过F12去查看，我们到底发送了什么包

![image-20230721185848408](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-211ad8-image-20230721185848408.png)

结果神奇的发现，这网站居然是用PHP写的。不过没关系，我们忽略掉图片之后，很容易就能找到这个网站发出的请求是什么，也就是一个拥有`{查询名字，大区编号，起始搜索位置，结束搜索位置，搜索类型（LOL/云顶），SC}`这六个字段的表单。

其中`SC`字段让我很不理解是什么，不过不着急，我们先用APIPost模拟一下请求。继续在浏览器中查看刚刚请求发送的方式，发现是`POST`

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-b1bcb2-image-20230721190151428.png)

然后我们就打开APIPost输入对应的字段:

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-ea9479-image-20230721190408426.png)

然后发现出错了：

![image-20230721190804931](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-00efac-image-20230721190804931.png)

看来`SC`字段是必传的，但是我们不大清楚`sc`到底是怎么来的。然后发现查询不同召唤师的`sc`还不同，不过查询同一个人的却是**相同**的。

### 7.24

![image-20230724111414268](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-5b58f6-image-20230724111414268.png)![image-20230724111346351](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-71ac67-image-20230724111346351.png)

看来`sc`应该是客户端自己生成的，因为我在最开始的页面刷新里并**没有看见`sc`从服务器**传过来。那就看看客户端的js是怎么生成`sc`的把。

可以发现它传了比较多的.js文件回来，我们一个个ctrl+f去找`sc`的变量名所在的js。

![image-20230724111853830](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-06565b-image-20230724111853830.png)

最后在game.js?v1.3.9中找到了`sc`的变量所在地方。

![image-20230724112103849](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-9b52e2-image-20230724112103849.png)

但是这个js文件中有很多奇奇怪怪的像16进制编码一样的东西，于是就跑去问chatgpt这是什么：

![](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-21-f7b243-image-20230721184544641.png)

坏消息是这个压缩后的变量名我根本看不懂，好消息是这是个变量名，我可以直接去找怎么生成的。于是就在这个网页里面找`_0x3848x20`的生成方式：

![image-20230724112409944](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-f12b5d-image-20230724112409944.png)

![image-20230724112454552](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-3fd836-image-20230724112454552.png)

可以发现`_0x3848x20`这个变量在上面几个函数里面都有用到，然后生成方式都是一样的，也就是：

``` js
 var _0x3848x20 = md5(__Ox101a53[0x24] + _0x3848x4 + __Ox101a53[0x25] + _0x3848x1f);
```

然后观察这个`function lolcx(_0x3848x4, _0x3848x1f) `函数，发现

![image-20230724164321186](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-ab941e-image-20230724164321186.png)

`name = _0x3848x4`，` dq = _0x3848x1f`，那么我们就已经找到两个参数了，接下来我们尝试去找`__Ox101a53`到底是什么。从ChatGPT了解到，这是一种JS混淆代码的技术。看上去我们无从下手，因为混淆之后的代码是不可能完全还原的，但这个的JS注释让我们知道了，他是在https://www.jsjiami.com进行的加密。

![image-20230724173257383](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-c788d5-image-20230724173257383.png)

于是我们顺藤摸瓜，找到了这个网站下的解密功能

![image-20230724173343599](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-a44edd-image-20230724173343599.png)

然后我们把解密的关键部分拷出来，也就是`__Ox101a53`，然后再根据`function lolcx()`的生成md5的逻辑看看到底丢进md5的字符串是什么：

``` js
var __Ox101a53 = ["getTime", "setTime", "cookie", "=", ";expires=", "toGMTString", "(^| )", "=([^;]*)(;|$)", "match", "data-lock", "attr", "#login", "true", "innerHTML", "loginload", "getElementById", "length", ".", "qrsig", "qrimg", "a", "querySelectorAll", "success", "ID\u5df2\u590d\u5236!", "msg", "on", "error", "options", "value", "selected", "substring", "search", "location", "&", "split", "\u6b63\u5728\u67e5\u8be2\u4e2d...", "name=", "dq-", "POST", "./api/nowapi.php?act=nowcx", "json", "ec", 
"show", "#dw1", "#dw2", "one", "html", "#one", "two", "#two", "#gamedata", "\u67e5\u8be2\u6210\u529f!", "timeout", "\u67e5\u8be2\u5931\u8d25!<br><br><br>\u591a\u6b21\u5931\u8d25\u8bf7\u52a0\u7fa4\u8054\u7cfb\u7fa4\u4e3b\u53cd\u9988", "alert", "\u670d\u52a1\u5668\u9519\u8bef\uff0c\u8bf7\u7a0d\u540e\u518d\u8bd5\uff01", "ajax", "#nowyx", '<a style="color:red;"href="/nowgame.php?name=', "&area=", '" target="_blank"> \u8be5\u53ec\u5524\u5e08\u6b63\u5728\u5bf9\u5c40\u4e2d.\u70b9\u51fb\u67e5\u770b\u8be6\u60c5</a>', 
"#nowyx1", "hide", "", "./api/nowapi.php?act=gg", "gg", "#gg", "right", "#right", "left", "#left", "name", "area", "val", "#name", "dq", ":checked", "is", "#ppp", "tft", "lol", "./api/lol.php?act=cx", "zhanji", "#sdata", "#log", ' |<span class="text-success" > Lv', "level", "</span>", "#sname", "privacy", "#syc", "lastGameDate", "#stime", "yy", '<tr height="50"><td onclick="send()" style="text-align:center;">\u4e0a\u4e00\u9875</td><td onclick="get_xyy(', "start", ",", "end", ');" style="text-align:center;">\u4e0b\u4e00\u9875</td></tr> ', 
"#yy", '<img class="img-avatar" src=https://wegame.gtimg.com/g.26-r.c2d3c/helper/lol/assis/images/resources/usericon/', "profileIconId", '.png alt=""></img>', "#sconid", "skin", "#skin", "#tftdw", "#ppdw", '<img src="//wegame.gtimg.com/g.26-r.c2d3c/helper/lol/v2/tier/tier-', "dstb", '.png" width="52" alt=""></img>', "#dstb", "dsdj", "#dsdj", "dssd", "#dssd", "dssf", "#dssf", "lhtb", "#lhtb", "lhdj", "#lhdj", "lhsd", "#lhsd", "lhsf", "#lhsf", "tfttb", "#tfttb", "tftdj", "#tftdj", "tftsd", "#tftsd", 
"tftsf", "#tftsf", "\u6700\u8fd1\u6ca1\u6709\u6bd4\u8d5b\u8bb0\u5f55\u4e86~", "\u5df2\u7ecf\u662f\u7b2c\u4e00\u9875\u4e86!", "#arealog", "#uin", "#nkey", "uinqqleader=", "\u8bf7\u8f93\u5165\u9700\u8981\u67e5\u8be2\u7684QQ", "\u8bf7\u8f93\u5165\u5361\u5bc6", "./api/lol.php?act=areaname", "#areaid", "\n", "\u67e5\u8be2\u5931\u8d25,\u8bf7\u91cd\u65b0\u5c1d\u8bd5<hr>\u591a\u6b21\u5931\u8d25\u8bf7\u52a0\u7fa4\u79c1\u804a\u7fa4\u4e3b", "#cdq option:selected", "#cname", "\u8bf7\u8f93\u5165\u9700\u8981\u67e5\u8be2\u7684ID", 
"999", "\u8bf7\u9009\u62e9ID\u5bf9\u5e94\u7684\u5927\u533a", "#uidlog", "#uid", "./api/lol.php?act=lolid", "#dq option:selected", '<tr height="50" ><td onclick="get_syy(', ');" style="text-align:center;">\u4e0a\u4e00\u9875</td><td onclick="get_xyy(', "\u6ca1\u6709\u66f4\u591a\u6218\u7ee9\u4e86!", '<tr height="50"><td onclick="send();" style="text-align:center;">\u4e0a\u4e00\u9875</td><td onclick="get_xyy(', ');" style="text-align:center;">\u4e0b\u4e00\u9875</tr> ', '<tr height="50"><td onclick="get_syy(', 
');" style="text-align:center;"> \u4e0b\u4e00\u9875</td></tr> ', "./api/lol.php?act=djxq", "win", "#dj", "loser", "loss", "time", "#djtime", "\u8bf7\u5148\u8f93\u5165\u9700\u8981\u67e5\u8be2\u7684ID", "\u8bf7\u9009\u62e9\u5927\u533a", "charCode", "which", "keyCode", "#000", "undefined", "log", "\u5220\u9664", "\u7248\u672c\u53f7\uff0cjs\u4f1a\u5b9a", "\u671f\u5f39\u7a97\uff0c", "\u8fd8\u8bf7\u652f\u6301\u6211\u4eec\u7684\u5de5\u4f5c", "jsjia", "mi.com"];

function gen(_0x3848x4, _0x3848x1f) {
    console.info(__Ox101a53[0x24] + _0x3848x4 + __Ox101a53[0x25] + _0x3848x1f)
}

gen("鸟倦飞知还", "1")
```

跑出来的结果是这个

![image-20230724173615701](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-d5d11a-image-20230724173615701.png)

然后根据查资料，md5如果字符一样的话，生成出来的md5的加密后的字符串也是一样的。于是我们去网上找一个32位md5加密的在线网站，把刚生成的md5字符丢进去：

![image-20230724173753955](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-f726fd-image-20230724173753955.png)

![image-20230724173826864](https://raw.githubusercontent.com/Vikyanite/talks/main/images/2023-07-24-47a3f5-image-20230724173826864.png)

**惊喜地发现，这就是我们想要！WOOOOOOWWWWWWOWOOWOWOW！**

特别感谢jsjiami.com！居然能解密，呜呜呜太激动了，还以为没办法了。

至此，我们总算可以进行下一步用代码模拟请求了！
