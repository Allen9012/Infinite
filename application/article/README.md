handler 转发层，参数校验和绑定
会调用logic 
业务逻辑都在logic层
svc主要是初始化一些第三方的依赖，初始化依赖于config
types 主要是


oss 主要内容是 上传文章的封面给到aliyun的oss 的bucket中
然后上传完之后需要给到数据库一个rpc请求去写入数据
那么就要初始化rpc  的客户端
然后就可以主要完成publish文章的操作了

下面开始处理article 的rpc 操作
首先编写一个proto文件
执行goctl rpc protoc ./article.proto --go_out=. --go-grpc_out=. --zrpc_out=./
然后生成表相关的内容： goctl model mysql datasource --dir ./internal/model --table article --url "root:9012@tcp(127.0.0.1:3306)/beyond_article"    