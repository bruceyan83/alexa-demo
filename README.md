#项目结构

```
+---cmd 可执行文件
|   \---lambda 该实例main函数所在位置
\---pkg
    +---alexa  alexa扩展
    +---lambda lambda扩展
    +---schema 数据结构验证
    \---virtual_device 虚拟设备状态控制，redis
```

```
cd cmd/lambda
GOOS=linux go build
zip handler.zip ./lambda
```

然后把这个zip传到您的lambda函数中，并且把lambda函数的handler名字改成lambda


### 需要您关注的环境变量
`REDIS_ADDR=xxx.xxx.xxx.xxx:6379;`\
`REDIS_PASSWORD=xxx;`

以上两条环境变量为配置virtual_device使用
