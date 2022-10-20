package main

import (
	"context"
	"log"
	"net"

	"go.unter.com/unter/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":7689"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}
	srv := grpc.NewServer()
	var u Unter
	pb.RegisterUnterServer(srv, &u)
	reflection.Register(srv)

	log.Printf("server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}

}
func (u *Unter) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// TODO: Validate Data
	log.Printf("Update: %v", req)
	resp := pb.UpdateResponse{
		CarId: req.CarId,
	}
	return &resp, nil
}

type Unter struct {
	pb.UnimplementedUnterServer
}
