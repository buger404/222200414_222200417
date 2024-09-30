## 参考 
[链接](https://gocn.github.io/styleguide/)

## 缩进
一般4个空格
## 变量命名
不用下划线，只在包内用的变量的首字母小写，如`newService`。
## 每行最多字符数
50
## 函数最大行数
不限
## 函数、类命名
同变量命名
## 常量
同变量命名
## 空行规则
不限，不要乱空行，比如返回一个结构体时，里面的元素排列就不应该有空行。

像下面这样
```
return &Watcher{
		watcher: w,
		files:   make(map[string]bool),
		done:    make(chan struct{}),
		dir:     make(map[string]bool),
	}, nil
```
不要写成这样
```
return &Watcher{

		watcher: w,
		files:   make(map[string]bool),
		done:    make(chan struct{}),
		dir:     make(map[string]bool),
	}, nil
```
## 注释规则
在函数头前注释，注释时带上函数名字。对你想注释的代码在前和右边注释都可，最好保持统一。
## 操作符前后空格
统一空一格