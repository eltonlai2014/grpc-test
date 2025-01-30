package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "example.com/project/mfa/proto" // 生成的 pb 代碼

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMfaServiceServer
}

func (s *server) GetQRcodeImg(ctx context.Context, req *pb.QRcodeRequest) (*pb.QRcodeResponse, error) {
	filename := req.GetMessage()
	filePath := fmt.Sprintf(".\\images\\%s", "qrcode.png")

	// 讀取圖片文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	log.Printf("Sending file: %s, size: %d bytes", filename, len(data))
	return &pb.QRcodeResponse{Filename: "rqcode.png", Data: data}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMfaServiceServer(s, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
