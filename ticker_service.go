package binance

import (
	"context"
	"encoding/json"
)

// ListBookTickersService list all book tickers
type ListBookTickersService struct {
	c *Client
}

// Do send request
func (s *ListBookTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*BookTicker, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ticker/allBookTickers",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*BookTicker, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// BookTicker define book ticker info
type BookTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

// ListPricesService list all ticker prices
type ListPricesService struct {
	c *Client
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ticker/allPrices",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*SymbolPrice, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// PriceChangeStatsService show stats of price change in last 24 hours
type PriceChangeStatsService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *PriceChangeStatsService) Symbol(symbol string) *PriceChangeStatsService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *PriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res *PriceChangeStats, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ticker/24hr",
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(PriceChangeStats)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}

	return
}

// PriceChangeStats define price change stats
type PriceChangeStats struct {
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FristID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type SymbolPriceService struct {
	c      *Client
	symbol string
}

// SymbolPrice define symbol and price pair
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// Symbol set symbol
func (s *SymbolPriceService) Symbol(symbol string) *SymbolPriceService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *SymbolPriceService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolPrice, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/ticker/price",
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(SymbolPrice)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}

	return
}
