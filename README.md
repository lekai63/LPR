# 个人记录


// TODO:删除银行借款合同中的current_lpr 并修改相应crud业务逻辑.因为已另外抓取lpr数据

## dlv远程调试

```
~/go/bin/dlv debug --headless --listen ":2345" --log --api-version 2
```

## use smallnest/gen to generate codes

```
~/go/bin/gen --sqltype=postgres \
   	--connstr "host=192.168.5.11 port=5432 user=fzzl dbname=lpr password=fzzl032003 sslmode=disable connect_timeout=10000" \
   	--database fzzl  \
   	--json \
   	--gorm \
   	--guregu \
   	--rest \
   	--out ./tables_gen \
   	--module github.com/lekai63/lpr/tables_gen \
   	--makefile \
   	--json-fmt=snake \
   	--overwrite
```

## postgres 操作

```
// 查看所有sequence
SELECT c.relname FROM pg_class c WHERE c.relkind = 'S'; 
// 修改序列起始号
alter sequence lease_contract_id_seq restart with 39;
// 删除记录
delete from xx_table where id = 9999 ;
```

## 调整时区
```
sudo timedatectl set-timezone 'Asia/Shanghai'
```

# GoAdmin 介绍

GoAdmin 是一个帮你快速搭建数据可视化管理应用平台的框架。 

- [github](https://github.com/GoAdminGroup/go-admin)
- [论坛](http://discuss.go-admin.com)
- [文档](https://book.go-admin.cn)

## 目录介绍

```
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile
├── adm.ini             adm配置文件
├── admin.db            sqlite数据库
├── build               二进制构建目标文件夹
├── config.json         配置文件
├── go.mod              go.mod
├── go.sum              go.sum
├── html                前端html文件
├── logs                日志
├── main.go             main.go
├── main_test.go        CI测试
├── pages               页面控制器
├── tables              数据模型
└── uploads             上传文件夹
```

## 生成CRUD数据模型

### 在线工具

管理员身份运行后，访问：http://127.0.0.1:8080/admin/info/generate/new

### 使用命令行工具

```
adm generate -l cn -c adm.ini
```

