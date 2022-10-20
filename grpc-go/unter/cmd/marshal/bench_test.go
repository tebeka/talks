package main

import (
	"encoding/json"
	"testing"

	"go.unter.com/unter/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	m = pb.UpdateRequest{
		Driver: "Bond",
		CarId:  "007",
		Location: &pb.Location{
			Lat: 1,
			Lng: 2,
		},
		Time:   timestamppb.Now(),
		Status: pb.Status_FREE,
	}
)

func BenchmarkProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&m)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&m)
		if err != nil {
			b.Fatal(err)
		}
	}
}
