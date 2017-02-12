# GetMeizi

[![Build Status](https://travis-ci.org/HackeZ/getMeizi.svg?branch=master)](https://travis-ci.org/HackeZ/getMeizi)

![image](https://github.com/HackeZ/getMeizi/blob/master/img/head.jpg)

同样是并发练手项目，我将使用 `gank.io` 的 [API接口](http://gank.io/api) 并发下载妹子图片

## Usage

```sh
> go get -u github.com/hackez/getmeizi...
> $GOPATH/bin/getMeizi --help # READ COMMAND HELP
```

## Dependence

- [kingpan](https://github.com/alecthomas/kingpin)

## Done

1. 指定下载路径完成。

1. 并发下载完成，我使用了一个Channel作为同步，达到了阻塞主进程的效果，但是我在代码中的表现方式并不尽人意，我会继续看看有没有其他更好的同步方法。

1. 通过select实现了 **MaxGORO** 限制了下载线程上限的功能，你只需要修改常量MaxGORO的值即可实现同时下载MaxGORO张图片的功能。

1. 现在通过2个Channel进行下载控制，Schedule表示下载线程线程池，Sign表示下载任务个数。

1. 通过正则表达式智能匹配到BaseURL中下载图片的张数，防止出现线程出现 **无限阻塞** 现象。

1. 添加 **响应超时提醒**

1. 将非缓存通道 sign 换成 WaitGroup

1. 将MaxGoro设置为 -1 ，即表示无限制Goroutine上限。

1. 将 flag 升级了可配置项

1. 将 flag 升级为 **kingpan**

## TODO

1. 按照妹子的姓名生成文件夹，然后将该妹子对应的图片存放起来。
1. 显示下载进程。


## Download Used Time(10 pictures)

1. 10.674s --> linear

1. 5.649s  --> set MaxGoro 3

1. 3.760s  --> MaxGoro be equal to Total Download Task

1. 5.709s  --> trans chan byte to WaitGroup

1. 2.991s  --> flag insert regexp

## Docker

[Docker Page](https://hub.docker.com/r/hackerz/getmeizi/)

## LICENSE

`MIT` :)
