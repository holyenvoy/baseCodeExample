package main

//client.go

import (
	"fmt"
	"log"
	"os"
	"time"

	pb "baseCodeExample/grpcHello/protobuf"

	"encoding/binary"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = ""
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	HelloRequest := &pb.HelloRequest{Type: 0, Name: name}

	PbBytes, err := proto.Marshal(HelloRequest)
	if err != nil {
		return
	}

	//fmt.Printf("PbBytes:%v size:%v\n", PbBytes, binary.Size(PbBytes))

	HelloRequest.Type = 0
	PbBytes, err = proto.Marshal(HelloRequest)
	if err != nil {
		return
	}
	start := time.Now()

	r, err := c.SayHello(context.Background(), HelloRequest)
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}

	fmt.Printf("duaration:%v\n", time.Since(start).Nanoseconds()/int64(time.Microsecond))

	fmt.Printf("PbBytes:%v, size:%v\n", PbBytes, binary.Size(PbBytes))
	log.Printf("Greeting: %s", r.Message)
}
