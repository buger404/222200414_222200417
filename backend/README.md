## 启动

### docker

切换到项目根目录，然后构建你的镜像

```
docker build -t your-image-name 
```

然后运行,这个前面的8888端口可以随便换

```
docker run -d -p 8888:8888 your-image-name
```

