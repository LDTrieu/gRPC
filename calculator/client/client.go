package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/ldtrieu/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)
	//log.Printf("service client %f", client)

	//callSum(client)
	//callPND(client)
	callAverage(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 6,
	})
	if err != nil {
		log.Fatalf("call sum api err %v", err)
	}
	log.Printf("sum api response %v \n", resp.GetResult())
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.PrimenumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err == io.EOF {
		log.Println("Server fisish streaming")
		return
	}
	if err != nil {
		log.Fatalf("callPND err %v", err)
	}
	for {

		resp, recErr := stream.Recv()
		if recErr == io.EOF {
			log.Println("Server finish streaming")
			return

		}
		log.Printf("prime number %v", resp.GetResult())

	}
}

func callAverage(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling ")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal("call averga err %v", err)
	}
	lisReq := []calculatorpb.AverageRequest{
		calculatorpb.AverageRequest{
			Num: 5,
		},
		calculatorpb.AverageRequest{
			Num: 10,
		},
		calculatorpb.AverageRequest{
			Num: 12,
		},
		calculatorpb.AverageRequest{
			Num: 20,
		},
		calculatorpb.AverageRequest{
			Num: 7,
		},
		calculatorpb.AverageRequest{
			Num: 2.3,
		},
	}
	for _, req := range lisReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("recive averge response err %v", err)

	}
	log.Printf("average response %+v", resp)
}
