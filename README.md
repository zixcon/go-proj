# go-proj
#### 特别声明：
#### 1. 项目中所有内容只为技术学习和分享，部分资源会进行摘录，如有侵权，请及时告知删除或者进行标示出处。
#### 2. 项目中所有内容在学习过程中请自觉遵守相关法律法规，由于学习本项目出现的任何问题均与本人无关，请自行处理！


#### 问题及解决
###### go get 出错
> package golang.org/x/net/html: unrecognized import path "golang.org/x/net/html" (https fetch: Get https://golang.org/x/net/html?go-get=1: net/http: TLS handshake timeout)

> 解决办法：
官网被墙，因此我们只能从github上拿到这部分包，放入项目中。以免下次go get的时候又去golang官网去找。
git clone  https://github.com/golang/net
在gopath目录简历如下目录  golang.org/x/net


#### 参考
> github.com/hunterhug/GoSpider
