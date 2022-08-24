# go-soar 介绍
一个基于 XiaoMi/soar 的 sql 优化工具，能批量给出优化建议

# 配置
1. 拉代码
2. 新增环境变量文件，并配置
```
cp .env.exam .env
```
3. 配置 mysql 连接和 soar 规则(重点)
修改 soar 下的 soar.yaml 文件

# 部署
1. build 镜像
```bash
docker build -t go-soar:v1 .
```
2. 起容器
```bash
docker run --name go-soar  -d  -p 3000:3000 go-soar:v1
```

# 使用
1. 获取单条 sql 优化建议
```bash
 curl -X POST localhost:3000/v1/soar/sql  -H 'Content-Type:multipart/form-data' -F 'sql=SELECT * FROM live_order WHERE user_id =11963232 ORDER BY order_number DESC'
```
2. 批量获取优化建议:上传需要分析的.sql 文件
```bash
curl -X POST localhost:3000/v1/soar/file  -H 'Content-Type:multipart/form-data' -F "sql_file=@E:\test.sql"
```

**返回的地址即为优化建议文件存放的地址。**
