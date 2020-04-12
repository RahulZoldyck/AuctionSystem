package api

import (
	"../model"
)

type AddBidderRequest struct {
	Port int    `json:"port"`
	Name string `json:"name"`
}
type AddBidderResponse struct {
	Code int `json:"code"`
}

type GetBidderListRequest struct {
}

type GetBidderListResponse struct {
	List []model.Bidder `json:"list"`
	Code int            `json:"code"`
}

type FindWinnerRequest struct {
	AuctionId string `json:"auction_id"`
}

type FindWinnerResponse struct {
	BidderId string  `json:"bidder_id"`
	BidValue float64 `json:"bid_value"`
	Code     int     `json:"code"`
}

type BidRequest struct {
	AuctionId string `json:"auction_id"`
}

type BidResponse struct {
	BidderId string  `json:"bidder_id"`
	BidValue float64 `json:"bid_value"`
	Code     int     `json:"code"`
}
