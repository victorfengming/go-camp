package main

import (
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 36000,
		FeeCent:     13000,
		StartPos: &trippb.Location{
			Latitude:30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:30,
			Longitude: 120,
		},
		PathLocaitons: []*trippb.Location{
			{
				Latitude:31,
				Longitude: 118,
			},
			{
				Latitude:37,
				Longitude: 125,
			},
		}, 
		Status:trippb.TripStaus_IN_PROGRESS,
	}

	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)
	var trip2 trippb.Trip
	proto.Unmarshal(b,&trip2)
	
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	bb,err:=json.Marshal(&trip2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("json:%s\n",bb)
}
