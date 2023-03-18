

# 测试

``` shell
# server
go run *.go -server=true -P=compact -buffered=true -framed=false -addr=localhost:9090 -secure=false
# client
go run *.go -server=false -P=compact -buffered=true -framed=false -addr=localhost:9090 -secure=false
```

