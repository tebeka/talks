package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.unter.org/unter/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	addr := ":7689"

	creds := insecure.NewCredentials()
	conn, err := grpc.DialContext(
		context.Background(), // if using WithBlock() - use a timeout here
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("connected to %s", addr)

	req := pb.UpdateRequest{
		CarId: "007",
		Location: &pb.Location{
			Lat: 1.1,
			Lng: 2.2,
		},
		Time: timestamppb.Now(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	c := pb.NewUnterClient(conn)
	resp, err := c.Update(ctx, &req)
	if err != nil {
		log.Fatalf("error: update - %s", err)
	}
	fmt.Println(resp)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	stream, err := c.Location(ctx)
	if err != nil {
		log.Fatalf("error: location - %s", err)
	}

	for i := 0; i < 5; i++ {
		req := pb.LocationRequest{
			CarId: "007",
			Location: &pb.Location{
				Lat: float64(i),
				Lng: float64(i),
			},
		}
		log.Printf("location send: %v", &req)
		if err := stream.Send(&req); err != nil {
			log.Fatalf("location send: %s", err)
		}
		time.Sleep(time.Second)
	}
	locResp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("location close: %s", err)
	}
	fmt.Println(locResp)
}
