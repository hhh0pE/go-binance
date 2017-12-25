package binance

import (
	"context"
	"encoding/json"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ping",
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	return
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	j, err := newJSON(data)
	if err != nil {
		return
	}
	serverTime = j.Get("serverTime").MustInt64()
	return
}

type ExchangeInfoService struct {
	c *Client
}

type SymbolUnit struct {
	Symbol             string `json:"symbol"`
	Status             string `json:"status"`
	BaseAsset          string `json:"baseAsset"`
	BaseAssetPrecision int    `json:"baseAssetPrecision"`
	QuoteAsset         string `json:"quoteAsset"`
	QuotePrecision     int    `json:"quotePrecision"`
}

type ExchangeInfo struct {
	Timezone   string       `json:"timezone"`
	ServerTime int64        `json:"serverTime"`
	Symbols    []SymbolUnit `json:"symbols"`
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/exchangeInfo",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}

	return
}
