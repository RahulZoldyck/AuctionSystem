package model

type Bid struct {
	Id    string  `json:"bid_id"`
	Value float64 `json:"bid_value"`
}
type Bids []Bid

func (b Bids) FindHighestBid() *Bid {
	var highest *Bid
	for _, v := range b {
		if highest == nil || highest.Value < v.Value {
			highest = &v
		}
	}
	return highest
}
