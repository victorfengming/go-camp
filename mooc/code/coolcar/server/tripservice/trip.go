package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

// 一定要加注释 注释要以TripService开头 .结尾

// Service impl a trip service.
type Service struct {
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 36000,
			FeeCent:     13000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			PathLocaitons: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 118,
				},
				{
					Latitude:  37,
					Longitude: 125,
				},
			},
			Status: trippb.TripStaus_IN_PROGRESS,
		},
	}, nil

}
