package hello_controller

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"hello/go_server/proto/hello"
)

type HelloController struct{}

func (h *HelloController) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	mp := propagation.MapCarrier{}
	for key, val := range md {
		mp[key] = val[0]
	}
	fmt.Println("---------------")
	fmt.Println(mp)
	fmt.Println("---------------")

	ctx = otel.GetTextMapPropagator().Extract(ctx, mp)

	return &hello.HelloResponse{Message: fmt.Sprintf("%s", in.Name)}, nil
}

func (h *HelloController) LotsOfReplies(in *hello.HelloRequest, stream hello.Hello_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&hello.HelloResponse{Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
	}
	return nil
}
