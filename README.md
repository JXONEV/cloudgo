#**服务计算cloudgo** 

实现了基本要求，通过了curl和ab测试。

-----

    实验选择使用了Martini框架，通过github上go-martini教程，对它有了初步的了解。

1.首先安装martini包，在src/github.com目录下安装

        go get github.com/go-martini/martini

2.安装完成后找到main.go的路径，执行

        go run main.go -p9090

得到网页和命令行的结果如下：

![](https://github.com/JXONEV/cloudgo/raw/master/image/1.png)

3.curl测试：使用两个cmd窗口，执行

        curl -v http://localhost:9090/hello/JXONEV
        
得到结果如下：

![](https://github.com/JXONEV/cloudgo/raw/master/image/2.png)

由于windows系统不自带curl命令，所以直接从网上下载并将CURL.EXE添加进windows/system32中，这样可以直接在命令行使用。

4.ab测试：在其中一个cmd上执行

        ab -n 1000 -c 100 http://localhost:9090/hello/world
        
得到相关信息反馈如下：

![](https://github.com/JXONEV/cloudgo/raw/master/image/3.png)

![](https://github.com/JXONEV/cloudgo/raw/master/image/4.png)

同样由于windows系统没有Apache的ab工具，所以从官网上下载了Apache服务器，从cmd中找到ab.exe的目录下进行的测试。

从上图可以得到的信息有：1000个请求全部成功，平均每个请求接收时间为1.09ms，传输速率为114.7kb/s，其中一半的请求在81ms内响应，所有请求完成相应需要443ms。
              
