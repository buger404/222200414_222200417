## 简介
本项目的后端主要功能为使用爬虫抓取巴黎奥运会的数据并返回给前端（爬虫仅用于作业，无恶意）。  

爬虫用的框架是`goquery`, http框架用的`hertz`(自动生成接口的功能比较好用)。

## 启动

### docker

切换到项目根目录，然后构建你的镜像

```
docker build -t your-image-name .
```

然后运行,这个前面的8888端口可以随便换

```
docker run -d -p 8888:8888 your-image-name
```

### Makefile

请先配好go的环境，然后输入`make run`, 关闭的时候输入`make stop`
