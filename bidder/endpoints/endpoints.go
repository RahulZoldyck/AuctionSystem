package endpoints

import (
	"../../api"
	"../service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeBidEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(api.BidRequest)
		id, value, err := svc.Bid(req.AuctionId)
		if err != nil {
			return api.BidResponse{BidderId: "nil", BidValue: -1}, err
		}
		return api.BidResponse{BidderId: id, BidValue: value, Code: 1}, nil
	}
}
