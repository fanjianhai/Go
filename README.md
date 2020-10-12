## 项目介绍

<div style="text-align: center;">
<img src="http://c.biancheng.net/uploads/allimg/180808/1-1PPQA9545W.jpg"/>
</div>
<br/>

**项目地址:** https://github.com/fanjianhai/Go
<br/><br/> go语言学习之路！

**持续更新...** 

## 项目结构

```
├── Go
│   ├── 00-基础语法
│      ├── 01-环境搭建&gomod依赖管理&goland第一个go程序
│      ├── 02-变量
│      ├── 03-数组
│      ├── 04-Slice 切片
│      ├── 05-Struct 结构体
│      ├── 06-Map 结构体
│      ├── 07-循环
│      ├── 08-函数
│      ├── 09-chan 通道
│      ├── 10-defer 函数
│      ├── 11-解析 JSON 数据
│      ├── 12-json.Unmarshal 遇到的小坑
│   ├── 01-Gin框架
│      ├── 01-框架安装
│      ├── 02-路由配置
│      ├── 03-使用 Logrus 进行日志记录
│      ├── 04-数据绑定和验证
│      ├── 05-自定义错误处理
│   ├── 02-gRPC
│   ├── 03-go-gin-api [文档]
├── Docker
├── Kubernetes
```

## 源码指引

#### Go - 基础篇

- :white_check_mark: [Go 基本语法](https://github.com/fanjianhai/Go/tree/master/00-%E5%9F%BA%E7%A1%80%E8%AF%AD%E6%B3%95/codes)

#### Gin 框架

- :white_check_mark: [Gin 框架 - 自定义路由配置](https://github.com/fanjianhai/Go/tree/master/01-Gin%E6%A1%86%E6%9E%B6/codes/02-%E8%B7%AF%E7%94%B1%E9%85%8D%E7%BD%AE)

- :white_check_mark: [Gin 框架 - 使用 Logrus 进行日志记录](https://github.com/fanjianhai/Go/tree/master/01-Gin%E6%A1%86%E6%9E%B6/codes/03-%E6%97%A5%E5%BF%97%E8%AE%B0%E5%BD%95)

- :white_check_mark: [Gin 框架 - 绑定数据和验证](https://github.com/fanjianhai/Go/tree/master/01-Gin%E6%A1%86%E6%9E%B6/codes/04-%E6%95%B0%E6%8D%AE%E9%AA%8C%E8%AF%81%E5%92%8C%E7%BB%91%E5%AE%9A)

- :white_check_mark: [Gin 框架 - 自定义错误处理](https://github.com/fanjianhai/Go/tree/master/01-Gin%E6%A1%86%E6%9E%B6/codes/05-%E8%87%AA%E5%AE%9A%E4%B9%89%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)


## github图片不显示问题
1. 源码中都含有相应的图片
2. 修改hosts文件
- 打开路径C:\Windows\System32\drivers\etc下的hosts文件,添加如下内容
```
# GitHub Start 
192.30.253.112    github.com 
192.30.253.119    gist.github.com
151.101.184.133    assets-cdn.github.com
151.101.184.133    raw.githubusercontent.com
151.101.184.133    gist.githubusercontent.com
151.101.184.133    cloud.githubusercontent.com
151.101.184.133    camo.githubusercontent.com
151.101.184.133    avatars0.githubusercontent.com
151.101.184.133    avatars1.githubusercontent.com
151.101.184.133    avatars2.githubusercontent.com
151.101.184.133    avatars3.githubusercontent.com
151.101.184.133    avatars4.githubusercontent.com
151.101.184.133    avatars5.githubusercontent.com
151.101.184.133    avatars6.githubusercontent.com
151.101.184.133    avatars7.githubusercontent.com
151.101.184.133    avatars8.githubusercontent.com
 
 # GitHub End
```
- 清空dns和浏览器缓存，dns缓存清空命令 ipconfig /flushdns
- 刷新浏览器

## 作者
昵称：雪山飞狐<br/>
格言：我可闭于一核桃壳内, 而仍自认我是无疆限之君主。<br/>
博客：https://blog.csdn.net/fanjianhai <br/>
邮箱：594042358@qq.com <br/>




