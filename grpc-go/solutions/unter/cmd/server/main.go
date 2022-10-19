package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"go.unter.org/unter/pb"
)

func main() {
	addr := ":7689"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}
	/* TLS
	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("error: can't find creds - %s", err)
	}
	srv := grpc.NewServer(grpc.Creds(creds))
	*/
	srv := grpc.NewServer(grpc.UnaryInterceptor(timeintInterceptor))
	var u Unter
	pb.RegisterUnterServer(srv, &u)
	reflection.Register(srv)

	log.Printf("server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func timeintInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if info.FullMethod != "/Unter/Update" {
		return handler(ctx, req)
	}

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("%s took %v", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}

func (u *Unter) Location(stream pb.Unter_LocationServer) error {
	var resp pb.LocationResponse

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Printf("location recv: %s", err)
			return err
		}

		log.Printf("location request: %v", req)
		resp.CarId = req.CarId // TODO: Check for mismatch
		resp.NumMessages++
	}

	if err := stream.SendAndClose(&resp); err != nil {
		return err
	}

	return nil
}

func validateRequest(r *pb.UpdateRequest) error {
	if r.CarId == "" {
		return fmt.Errorf("empty car ID")
	}
	// TODO: more validations

	return nil
}

func (u *Unter) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// TODO: Validate
	if err := validateRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	log.Printf("Update: %v", req)
	resp := pb.UpdateResponse{
		CarId: req.CarId,
	}

	return &resp, nil
}

type Unter struct {
	pb.UnimplementedUnterServer
}
