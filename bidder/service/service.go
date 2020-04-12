package service

import (
	"os"
	"strconv"
	"time"
)

type Service interface {
	Bid(auctionID string) (string, float64, error)
}

type BiddingService struct {
	Name  string
	Delay int
}

func (b *BiddingService) Bid(auctionID string) (string, float64, error) {
	time.Sleep(time.Duration(b.Delay) * time.Millisecond)
	value, err := strconv.ParseFloat(os.Getenv("BIDDER_VALUE"), 64)
	return b.Name + "-" + auctionID, value, err
}
