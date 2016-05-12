#GetMeizi#

![image](https://github.com/HackeZ/getMeizi/blob/master/img/head.jpg)

###同样是并发练手项目，我将使用gank.io的API接口并发下载妹子图片###

###Use Example###

> go build -o getMeizi main.go

> ./getMeizi /**download path**/

===


###本项目同时适用于Windows以及Linux平台，如果你在Windows平台中出现下载错误，请关闭包括360安全卫士、百度卫士、腾讯管家在内等多家的流氓软件！###

> P.S. 因为Windows以及Linux操作系统不同，指定的download path请自行参考对应的操作系统命令

###Next Step###

###指定下载路径完成。###

###并发下载完成，我使用了一个Channel作为同步，达到了阻塞主进程的效果，但是我在代码中的表现方式并不尽人意，我会继续看看有没有其他更好的同步方法。###

###通过select实现了 **MaxGORO** 限制了下载线程上限的功能，你只需要修改常量MaxGORO的值即可实现同时下载MaxGORO张图片的功能。###

###现在通过2个Channel进行下载控制，Schedule表示下载线程线程池，Sign表示下载任务个数###

###通过正则表达式智能匹配到BaseURL中下载图片的张数，防止出现线程出现 **无限阻塞** 现象###

###接下来我想做的有：###

> 1. 按照妹子的姓名生成文件夹，然后将该妹子对应的图片存放起来。


**Download Used Time**

> 1. 10.674s

> 2. 5.649s

> 3. 3.760s
