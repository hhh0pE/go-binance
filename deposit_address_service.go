package binance

import (
	"context"
	"encoding/json"
)

// ListDepositsService list deposits
type DepositAddressService struct {
	c         *Client
	asset     *string
	timestamp *int64
}

// Asset set asset
func (s *DepositAddressService) Asset(asset string) *DepositAddressService {
	s.asset = &asset
	return s
}

// StartTime set timetsamp
func (s *DepositAddressService) Timestamp(startTime int64) *DepositAddressService {
	s.timestamp = &startTime
	return s
}

// Do send request
func (s *DepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *DepositAddressResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/depositAddress.html",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.asset != nil {
		m["asset"] = *s.asset
	}
	if s.timestamp != nil {
		m["timestamp"] = *s.timestamp
	}

	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DepositHistoryResponse define deposit history
type DepositAddressResponse struct {
	Success  bool       `json:"success"`
	Address string `json:"address"`
	AddressTag string `json:"addressTag"`
	Asset string `json:"asset"`
}