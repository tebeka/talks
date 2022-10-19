package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.unter.org/unter/pb"
)

func main() {
	loc := pb.Location{
		Lat: -22.95,
		Lng: -43.21,
	}
	fmt.Println(&loc)
	/* Copy of sync.Mutex - don't do that - work with pointers
	loc2 := loc
	fmt.Println(loc2)
	*/

	data, err := proto.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("proto size:", len(data))
	jdata, err := json.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json size:", len(jdata))

	var loc2 pb.Location
	if err := proto.Unmarshal(data, &loc2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("loc2:", &loc2)

	req := pb.UpdateRequest{
		CarId:  "007",
		Driver: "Bond",
		Location: &pb.Location{
			Lat: -22.95,
			Lng: -43.21,
		},
		Time:   timestamppb.Now(),
		Status: pb.Status_RIDING,
	}
	fmt.Println(&req)
	t := req.Time.AsTime() // protobuf -> Go
	fmt.Println("go:", t)
	pt := timestamppb.New(t) // Go -> protobuf
	fmt.Println("pb:", pt)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(&req)
	enc.Encode(t)

	data, err = protojson.Marshal(&req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
