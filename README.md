## 目前包含27个app，不断添加中
## 代码很简单粗暴，以后再精简优化

+ 360搜索
+ 哔哩哔哩
+ AcFun
+ CSDN
+ 懂球帝
+ 豆瓣
+ 抖音
+ GitHub
+ 国家地理
+ 历史上的今天
+ 虎扑
+ IT之家
+ 梨视频
+ 澎湃新闻
+ 腾讯新闻
+ 少数派
+ 搜狗
+ 今日头条
+ V2EX
+ 网易新闻
+ 微博
+ 新京报
+ 知乎
+ 夸克
+ 搜狐
+ 百度
+ 南方周末

## 运行

`go run main.go`

默认端口为1111，可以自己改。

浏览器打开`ip:1111`

具体路径看`main.go`中的路由表

比如访问百度的热搜就是：`ip:1111/baidu`

查看全部app的热搜为`ip:1111/all`

建议`go build`为可执行文件。