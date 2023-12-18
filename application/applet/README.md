applet是一个bff的服务
bff主要负责响应和请求的封装（http）
是一个http的服务


applet的api参考api的规范
https://go-zero.dev/docs/tutorials
首先按照要求定义applet.api

go-zero对应api代码生成
```bash
goctl api go --dir=./ --api applet.api
```
