用户服务

创建rpc目录，在目录下实现user.proto

自动注册中心，查看etcd中是否已经注册
```bash
docker exec eager_jones etcdctl get --prefix user.rpc
# 查看是否还有租约
docker exec eager_jones etcdctl lease list
# 查看租约剩余时间
docker exec eager_jones etcdctl lease timetolive 
```