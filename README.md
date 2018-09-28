# apidocserver

<code>golang 1.10</code> &nbsp; <code>echo 3.3</code> &nbsp; <code>redis 3.0.5</code> &nbsp; <code>mysql 8.0</code>

#### 测试地址：
                 http://apidoc.amagic.top

#### 账号密码：
                admin
                123456


#### Build Setup
                本项目用redis保存了用户的登录token
                本项目采用MySql保存基础数据
                你可以在conf.yaml文件中配置redis和MySql的基础信息


                go build app.go

#### 需要下载库
               web框架：go get github.com/labstack/echo
               xorm：go get github.com/go-xorm/xorm
               mysql驱动：go get github.com/go-sql-driver/mysql
               redis库：go get github.com/garyburd/redigo/redis
               websocket库：go get github.com/gorilla/websocket
               yaml解析库：go get github.com/kylelemons/go-gypsy/yaml



#### Contact me:
* Author : [@aixinaxc][1]
* Blog：[aixianxc-blog][2]

[1]: http://www.amagic.top/
[2]: https://blog.csdn.net/aixinaxc/
