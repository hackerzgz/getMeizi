#GetMeizi#

![image](https://github.com/HackeZ/getMeizi/blob/master/img/head.jpg)

###同样是并发练手项目，我将使用gank.io的API接口并发下载妹子图片##

###因为本项目在Windows平台下出现了不可预知的错误，所以先将项目放上来，逐渐解决问题，同样欢迎大家能够帮我看看怎么解决错误。谢谢！###

- 因为gank.io图片使用了sinaimg作为图床，但是我在用Golang访问sinaimg端口时候会出现以下错误：

> write tcp 10.15.32.248:63418->112.90.6.238:80: wsasend: An operation was attempted on something that is not a socket.

###或者是###

> read tcp 10.15.32.248:41454->112.90.152.14:80: wsarecv: An established connection was aborted by the software in your host machine.

###这样的错误只会出现在Windows平台中，在Linux中使用本项目是完全没有问题的，我通过百度，好像没有找到相关解决问题的答案，希望知道问题出现的朋友能够告诉我原因，感谢！###

###Next Step###

###指定下载路径完成###

###并发下载完成，我使用了一个Channel作为同步，达到了阻塞主进程的效果，但是我在代码中的表现方式并不尽人意，我会继续看看有没有其他更好的同步方法。###
