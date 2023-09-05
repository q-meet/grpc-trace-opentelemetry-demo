## 项目介绍

**目录结构**

```
├─ hello  -- 代码根目录
│  ├─ go_client
│     ├── main.go
│     ├── proto
│         ├── hello
│            ├── hello.pb.go
│  ├─ go_server
│     ├── main.go
│     ├── controller
│         ├── hello_controller
│            ├── hello_server.go
│     ├── proto
│         ├── hello
│            ├── hello.pb.go
│            ├── hello.proto
```

**hello.proto**

```
syntax = "proto3"; // 指定 proto 版本

package hello;     // 指定包名

// 定义 Hello 服务
service Hello {

	// 定义 SayHello 方法
	rpc SayHello(HelloRequest) returns (HelloResponse) {}

	// 定义 LotsOfReplies 方法
	rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}
}

// HelloRequest 请求结构
message HelloRequest {
	string name = 1;
}

// HelloResponse 响应结构
message HelloResponse {
    string message = 1;
}
```

## 效果

```
go run main.go

----------------
----------------traceIdStr: 00748bd593870f8202f660f54a880a37 <nil>
----------------spanIDStr: 3e1fe2f095a57398 <nil>
----------------
2023/09/06 00:37:53 Hello World
2023/09/06 00:37:53 Hello World Reply 0
2023/09/06 00:37:53 Hello World Reply 1
2023/09/06 00:37:53 Hello World Reply 2
2023/09/06 00:37:53 Hello World Reply 3
2023/09/06 00:37:53 Hello World Reply 4
2023/09/06 00:37:53 Hello World Reply 5
2023/09/06 00:37:53 Hello World Reply 6
2023/09/06 00:37:53 Hello World Reply 7
2023/09/06 00:37:53 Hello World Reply 8
2023/09/06 00:37:53 Hello World Reply 9



server:
---------------
map[:authority:0.0.0.0:9090 content-type:application/grpc traceparent:00-00748bd593870f8202f660f54a880a37-3e1fe2f095a57398-01 tracestate:foo=bar user-agent:grpc-go/1.57.0 xxx:222]
---------------

```
