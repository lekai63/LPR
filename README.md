# 个人记录

// TODO:错误梳理、处理

// 测试后修订main.go

# pgsql 相关

## postgres 操作

```
// 查看所有sequence
SELECT c.relname FROM pg_class c WHERE c.relkind = 'S'; 
// 修改序列起始号
alter sequence lease_contract_id_seq restart with 39;
// 删除记录
delete from xx_table where id = 9999 ;
```

## pgSQL在配置更新后的热重启

Command below is the perfect command to avoid restart and most importantly does not disrupt ongoing queries

```
docker exec -it {postgres_container}  psql -U postgres -c "SELECT pg_reload_conf();"
```

## 调整时区

```
sudo timedatectl set-timezone 'Asia/Shanghai'
```

# go开发相关

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

# 项目实践中参考使用了以下项目，致谢：
https://github.com/GoAdminGroup/go-admin
https://github.com/smallnest/gen
https://gorm.io/



