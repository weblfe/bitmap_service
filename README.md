## bitmap 服务

开发框架 go-zero(rpc,api,model,cache,mirco-service) + RoaringBitmap/roaring (压缩bitmap)

## 服务功能介绍

> 简介 

 提供 用户留存率计算服务

```
次留存   :（第一天新增用户 && 到第2天还登录的用户数）/  第一天总新增用户数 *100%
7日留存  :（第一天新增用户 && 到第8天还登录的用户数）/  第一天总新增用户数 *100%
15日留存 :（第一天新增用户 && 到第16天还登录的用户数）/ 第一天总新增用户数 *100%
30日留存 :（第一天新增用户 && 到第31天还登录的用户数）/ 第一天总新增用户数 *100%
```

> ### 获取单天 留存率

```http request
http://localhost:8080/bitmap/userkeeper?date=2021-02-01&day=1
```

渠道留存
```http request
http://localhost:8080/bitmap/userkeeper?date=2021-02-01&day=1&channel=xxx
```

> ### 获取多天留存率

date : 查询日期 
days : 距离留存起始日天数集合
channel : 渠道名
role : 身份类型
type : 查询日期类型 (0:截止日期 默认为 0,1:开始日期)

```http request
http://localhost:8080/bitmap/userkeepers?date=2021-02-01&days=[1,2,3,15,30]
```

渠道留存
```http request
http://localhost:8080/bitmap/userkeepers?date=2021-02-01&days=[1,2,3,15,30]&channel=xxx
```

多渠道多天留存
```http request
http://localhost:8080/bitmap/userkeepers?date=2021-02-01&days=[1,2,3,15,30]&channels=["debug","test"]
```

按身份留存
```http request
http://localhost:8080/bitmap/userkeepers?date=2021-02-01&days=[1,2,3,15,30]&role=1
```


## 服务启动命令

```bash
# 启动开发api服务
run.sh dev 
# 更新 协议生成代码
run.sh update

# 以 docker 方式 服务
run.sh stop
# 停止 docker 服务
run.sh stop
```

# Changelog

### v1.0.0

支持 用户留存 单天留存查询

支持 用户留存 多天查询

### v1.0.1

支持 用户 渠道留存(多天|单天)

### v1.0.2

支持 360天 留存查询

### v1.0.3

添加 多渠道 多天 留存查询

### v1.0.4

添加 按用户身份 多天 留存查询

### v1.0.5

添加 查询日期 类型(1:开始,0:截止)支持  查询