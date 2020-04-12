package main

import (
	"../api"
	"./endpoints"
	"./service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handleErr(err error, errString string) {
	if err != nil {
		if errString != "" {
			fmt.Println("Delay should be an integer (ms)")
		} else {
			fmt.Println(err.Error())
		}
		return
	}
}

func main() {
	bidderName := os.Getenv("BIDDER_NAME")
	delay := os.Getenv("BIDDER_DELAY")
	port := os.Getenv("BIDDER_PORT")
	delayInt, err := strconv.Atoi(delay)
	handleErr(err, "Delay should be an integer (ms)")

	portInt, err := strconv.Atoi(delay)
	handleErr(err, "Port should be an integer")

	err = registerWithAuctioneer(bidderName, portInt)
	handleErr(err, "")

	svc := service.BiddingService{Name: bidderName, Delay: delayInt}
	bidHandler := httptransport.NewServer(endpoints.MakeBidEndpoint(&svc), decodeBidRequest, encodeResponse)
	http.Handle("/bid", bidHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func registerWithAuctioneer(name string, port int) error {
	reqBody, err := json.Marshal(api.AddBidderRequest{Name: name, Port: port})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:8888", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()

		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var body api.AddBidderResponse
		err = json.Unmarshal(bodyByte, body)
		if err != nil {
			return err
		}
		if body.Code == 1 {
			fmt.Println("Successfully Registered")
			return nil
		} else {
			fmt.Println("Registration Failed")
			return errors.New("registration failed")
		}
	} else {
		fmt.Println("Registration Failed")
		return errors.New("registration failed")
	}
}

func decodeBidRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.BidRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
