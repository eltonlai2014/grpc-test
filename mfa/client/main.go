package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "example.com/project/mfa/proto" // 生成的 pb 代碼
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMfaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetQRcodeImg(ctx, &pb.QRcodeRequest{Message: "中文World"})
	if err != nil {
		log.Fatalf("Error while GetQRcodeImg: %v", err)
	}
	filename := resp.GetFilename()
	log.Printf("ImgName: %s", filename)

	// 保存文件
	outputPath := fmt.Sprintf(".\\%s", filename)
	err = os.WriteFile(outputPath, resp.GetData(), os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	fmt.Printf("Image saved to: %s\n", outputPath)
}
