package endpoints

import (
	"../../api"
	"../service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeAddBidderEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AddBidderRequest)
		v, err := svc.AddBidder(req.Name, req.Port)
		if err != nil {
			return api.AddBidderResponse{Code: 0}, err
		}
		return api.AddBidderResponse{Code: v}, nil
	}
}

func MakeGetBidderListEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(api.GetBidderListRequest)
		v, err := svc.GetBidderList()
		if err != nil {
			return api.GetBidderListResponse{List: nil, Code: 0}, err
		}
		return api.GetBidderListResponse{List: v, Code: 1}, nil
	}
}

func MakeFindWinnerEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(api.FindWinnerRequest)
		var id, value, err = svc.FindWinner(req.AuctionId)
		if err != nil {
			return api.FindWinnerResponse{BidderId: "nil", BidValue: -1}, err
		}
		return api.FindWinnerResponse{BidderId: id, BidValue: value, Code: 1}, nil
	}
}
