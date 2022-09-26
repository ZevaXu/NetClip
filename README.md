# NetClip

## 介绍
用Go语言开发的在线剪贴板项目，后端使用[Gin](https://gin-gonic.com)框架，前端使用[Bootsrap](https://getbootstrap.com)框架(之前没学过前端，本项目的页面是现学现写的),数据库使用[Mysql](https://www.mysql.com/)，采用[Gorm](https://gorm.io)框架与数据库交互。
本项目为学习Go语言的练手项目，同时也是自己一直有在线剪贴板这样一个需求的简单实现。

## 主要功能
提供一个在线的文本编辑、文本分享和文本暂存服务的web小工具，实现跨平台、跨设备的文本分享和暂存服务。
开发原因：由于自己有多个设备(IOS、安卓、mac、windows)，有时候想从手机发一个URL到windows电脑上打开，但又不想登录微信获取QQ等，然而不用这些软件又实在难以传输。。。。

## Demo
https://www.netclip.site

## 部分截图
1. 剪贴板创建页面
![20220926-2310-7rIeZI|550](https://cdn.jsdelivr.net/gh/ZevaXu/picupload@master/uPic/20220926-2310-7rIeZI%7C550.png)
2. 剪贴板编辑页面
![20220926-2309-xtAbb6|550](https://cdn.jsdelivr.net/gh/ZevaXu/picupload@master/uPic/20220926-2309-xtAbb6%7C550.png)

## 部署
1. 修改.env文件，添加数据库连接信息。
2. 使用docker-compose.yml文件和Dockerfile文件部署项目。

