package main

import (
	"context"
	"io"
	"log"
	"time"

	//"github.com/ldtrieu/grpc"

	"grpc/calculator/calculatorpb"

	//"github.com/grpc/grpc-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"

	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/internal/credentials"
	"google.golang.org/grpc/status"
)

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
func callFindMax(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling find max ...")
	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("call find max err %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		//gui nhieu request
		listReq := []calculatorpb.FindMaxRequest{
			calculatorpb.FindMaxRequest{
				Num: 5,
			},
			calculatorpb.FindMaxRequest{
				Num: 10,
			},
			calculatorpb.FindMaxRequest{
				Num: 12,
			},
			calculatorpb.FindMaxRequest{
				Num: 3,
			},
			calculatorpb.FindMaxRequest{
				Num: 4,
			},
		}
		for _, req := range listReq {
			err := stream.Send(&req)
			if err != nil {
				log.Fatalf("send find max request err %v", err)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("ending find max api ...")
				break
			}
			if err != nil {
				log.Fatalf("recv find max err %v", err)
				break
			}

			log.Printf("max: %v\n", resp.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}

func callSquareRoot(c calculatorpb.CalculatorServiceClient, num int32) {
	log.Println("calling square root api")
	resp, err := c.Square(context.Background(), &calculatorpb.SquareRequest{
		Num: num,
	})
	if err != nil {
		log.Printf("call square root api err %v", err)
		if errStatus, ok := status.FromError(err); ok {
			log.Printf("err msg: %v\n", errStatus.Message())
			log.Printf("err code: %v\n", errStatus.Code())
			if errStatus.Code() == codes.InvalidArgument {
				log.Printf("InvalidArgument num %v", num)
				return

			}
		}
	}
	log.Printf("sum square root response %v\n", resp.GetSquareRoot())
}

func callSumWithDeadline(c calculatorpb.CalculatorServiceClient, timeout time.Duration) {
	log.Println("calling sum with deadline api")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := c.SumWithDeadline(ctx, &calculatorpb.SumRequest{
		Num1: 7,
		Num2: 6,
	})
	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("calling sum with deadline DeadlineExceeded")
			} else {
				log.Fatalf("call sum with deadline api err %v", err)
			}
		} else {

			log.Fatalf("call sum with deadline unknown err %v", err)
		}
	}
	log.Printf("sum with deadline api response %v\n", resp.GetResult())
}

func main() {
	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("create client creds ssl err %v\n", sslErr)
		return
	}
	cc, err := grpc.Dial("localhost:50069", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)
	//log.Printf("service client %f", client)

	callSum(client)
	//callPND(client)
	//callAverage(client)
	//callSquareRoot(client, -4)
	//callSumWithDeadline(client, 1*time.Second)
	//callSumWithDeadline(client, 5*time.Second)
}
