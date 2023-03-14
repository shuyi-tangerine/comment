

# 测试

``` shell
# server
go run *.go -server=true -P=json -transport=buffered -addr=localhost:9090 -secure=false
# client
go run *.go -server=false -P=json -transport=buffered -addr=localhost:9090 -secure=false
```

