package main

import (
	"../api"
	"./endpoints"
	"./service"
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	svc := service.AuctionService{}
	addBidderHandler := httptransport.NewServer(endpoints.MakeAddBidderEndpoint(&svc), decodeAddBidderRequest, encodeResponse)
	getBidderListHandler := httptransport.NewServer(endpoints.MakeGetBidderListEndpoint(&svc), decodeGetBidderListRequest, encodeResponse)
	findWinnerHandler := httptransport.NewServer(endpoints.MakeFindWinnerEndpoint(&svc), decodeFindWinnerRequest, encodeResponse)
	http.Handle("/addbidder", addBidderHandler)
	http.Handle("/getbidderlist", getBidderListHandler)
	http.Handle("/findwinner", findWinnerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func decodeAddBidderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.AddBidderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeGetBidderListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.GetBidderListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeFindWinnerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.FindWinnerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
