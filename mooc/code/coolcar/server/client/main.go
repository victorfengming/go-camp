package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)

	}

	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "tr35653",
	})
	if err != nil {
		log.Fatalf("can not call GetTrip : %v", err)
	}

	fmt.Println(r)
}

/**
id:"tr35653"  trip:{start:"abc"  start_pos:{latitude:30  longitude:120}  end_pos:{latitude:30  longitude:120}  path_locaitons:{latitude:31  longitude:118}  path_locaitons:{latitude:37  longitude:125}  status:IN_PROGRESS  end:"def"  duration_sec:36000  fee_cent:13000}

Process finished with the exit code 0
*/
