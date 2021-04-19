package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc/com.deali/grpc"
	"io"
	"log"
	"net"
)

const (
	port = ":8888"
)

// grpc impl
type service struct {
	pb.GrpcServiceServer
}

func (s * service) GetOne(ctx context.Context,req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: []float64{1.1, 2.2, 3.2}}, nil
}

func (s * service) ServerStream(req *pb.Request, stream pb.GrpcService_ServerStreamServer) error {

	// 1요청, n응답
	stream.Send(&pb.Response{Value: []float64{1.1, 2.2, 3.2}})
	stream.Send(&pb.Response{Value: []float64{4.1, 5.2, 6.2}})
	stream.Send(&pb.Response{Value: []float64{7.1, 8.2, 9.2}})

	return nil
}

func (s * service) ClientStream(stream pb.GrpcService_ClientStreamServer) error {

	// n요청, 1응답
	for {
		req ,err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{Value: []float64{4.1, 5.2, 6.2}})
		}
		if err != nil {
			return err
		}
		log.Print(req)
	}
}

func (s * service) BiStream(stream pb.GrpcService_BiStreamServer) error {

	// n요청, n응답
	for {
		req, err := stream.Recv()
		log.Print(req)
		if err == io.EOF {
			return nil
		}
		for range make([]int, 5) {
			stream.Send(&pb.Response{Value: []float64{4.1, 5.2, 6.2}})
		}
	}
}

// grpc config
func main() {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen #{err}")
	}

	server := grpc.NewServer()
	pb.RegisterGrpcServiceServer(server, &service{})
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}