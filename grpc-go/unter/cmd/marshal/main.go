package main

import (
	"fmt"
	"os"

	"go.unter.com/unter/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	m := pb.UpdateRequest{
		Driver: "Bond",
		CarId:  "007",
		Location: &pb.Location{
			Lat: 1,
			Lng: 2,
		},
		Time:   timestamppb.Now(),
		Status: pb.Status_FREE,
	}
	// m := pb.UpdateRequest{} // 0 bytes

	data, err := proto.Marshal(&m)
	if err != nil {
		fmt.Printf("error: can't proto marshal - %s", err)
		return
	}
	fmt.Println("proto size:", len(data))

	// jdata, err := json.Marshal(&m)
	jdata, err := protojson.Marshal(&m)
	if err != nil {
		fmt.Printf("error: can't json marshal - %s", err)
		return
	}
	fmt.Println("json size: ", len(jdata))
	os.Stdout.Write(jdata)
}
