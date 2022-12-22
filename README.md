# go-cloud-storage

> 基于go-zero、xorm实现的云盘文件存储系统

### 配置
##### Mac/Linux
```shell
export MailFrom=""
export MailTo=""
export Mail163Pass=""
export MailServer=""
export MailServerPort=""

export QiNiuAK=""
export QiNiuSk=""
export QiNiuBucket=""
export QiuNiuUrl=""
```

##### Windows
```cmd
go env -w MailFrom=""
go env -w MailTo=""
go env -w Mail163Pass=""
go env -w MailServer=""
go env -w MailServerPort=""

go env -w QiNiuAK=""
go env -w QiNiuSk=""
go env -w QiNiuBucket=""
go env -w QiuNiuUrl=""
```


```shell
goctl api new server

goctl api go -api server.api -dir . -style go_zero
```


## 功能
- [x] 用户
    - [x] 密码登录
    - [x] 刷新token
    - [x] 邮箱注册
    - [x] 用户详情
- [x] 存储
    - [x] 中心存储资源管理
        - [x] 文件上传
        - [x] 文件秒传
        - [x] 文件分片上传
        - [x] 对接七牛云对象存储
    - [x] 个人存储资源管理
        - [x] 文件关联存储
        - [x] 文件列表
        - [x] 文件名称修改
        - [x] 文件夹创建
        - [x] 文件删除
        - [x] 文件移动
- [x] 文件分享
    - [x] 创建分享记录
    - [x] 获取资源详情
    - [x] 资源保存
    