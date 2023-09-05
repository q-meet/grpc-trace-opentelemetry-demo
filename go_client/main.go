package main

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hello/go_client/proto/hello"
	"io"
	"log"
	"math/rand"
	"time"
)

const (
	// gRPC 服务地址
	Address = "0.0.0.0:9090"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	var ctx = context.Background()
	ctx = context.WithValue(ctx, "xx", "xx1123")
	TraceState := trace.TraceState{}
	TraceState, _ = TraceState.Insert("foo", "bar")
	traceIdStr := generateTraceID(32)
	spanIDStr := generateTraceID(16)
	traceId, err := trace.TraceIDFromHex(traceIdStr)
	spanID, err := trace.SpanIDFromHex(spanIDStr)
	fmt.Println("----------------")
	fmt.Println("----------------traceIdStr:", traceIdStr, err)
	fmt.Println("----------------spanIDStr:", spanIDStr, err)
	fmt.Println("----------------")
	span := trace.SpanContextConfig{
		TraceID:    traceId,
		SpanID:     spanID,
		TraceFlags: 0x1,
		TraceState: TraceState,
		Remote:     true,
	}
	spanCtx := trace.NewSpanContext(span)

	ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)

	md := propagation.MapCarrier{
		"xxx": "222",
	}
	otel.GetTextMapPropagator().Inject(ctx, md)
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(md))

	// 调用 SayHello 方法
	res, err := c.SayHello(ctx, &hello.HelloRequest{Name: "Hello World"})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)

	// 调用 LotsOfReplies 方法
	stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}

		log.Printf("%s", res.Message)
	}
}

func generateTraceID(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdef01234567890123456789"
	//6261559c186f37f6f8e7018615569d1d
	traceID := make([]byte, length)
	for i := 0; i < length; i++ {
		traceID[i] = charset[rand.Intn(len(charset))]
	}

	return string(traceID)
}
