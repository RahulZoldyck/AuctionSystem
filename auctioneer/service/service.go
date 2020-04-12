package service

import (
	"../../api"
	"../../model"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Service interface {
	AddBidder(name string, port int) (int, error)
	GetBidderList() ([]model.Bidder, error)
	FindWinner(auctionID string) (string, float64, error)
}

type AuctionService struct{
	bidders []model.Bidder
}

func (a AuctionService) AddBidder(name string, port int) (int, error) {
	a.bidders = append(a.bidders, model.Bidder{Port: port, Name: name})
	return 1, nil
}
func (a AuctionService) GetBidderList() ([]model.Bidder, error) {
	return a.bidders, nil
}
func (a AuctionService) FindWinner(auctionID string) (string, float64, error) {
	ch := make(chan model.Bid, len(a.bidders))
	errCh := make(chan error)
	for _, v := range a.bidders {
		go func() {
			reqBody, err := json.Marshal(api.BidRequest{AuctionId: auctionID})
			if err != nil {
				errCh <- err
			}
			req, err := http.NewRequest("POST", "http://localhost:"+string(v.Port), bytes.NewBuffer(reqBody))
			if err != nil {
				errCh <- err
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				errCh <- err
			}
			if resp != nil && resp.Body != nil {
				defer resp.Body.Close()

				bodyByte, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					errCh <- err
				}
				var body api.BidResponse
				err = json.Unmarshal(bodyByte, body)
				if err != nil {
					errCh <- err
				}
				ch <- model.Bid{Id: body.BidderId, Value: body.BidValue}

			}
		}()

	}

	var bids model.Bids

	for {
		select {
		case r := <-ch:
			bids = append(bids, r)
		case e := <-errCh:
			return "nil", -1, e
		case <-time.After(200 * time.Millisecond):
			highestBid := bids.FindHighestBid()
			if highestBid == nil {
				return "nil", -1, errors.New("no bid found")
			}
			return highestBid.Id, highestBid.Value, nil
		}
	}

}


