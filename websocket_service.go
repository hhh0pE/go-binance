package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/hhh0pE/decimal"
	"github.com/hhh0pE/golang-json-decoder-case-sensitive"
	"go.uber.org/zap"
)

var (
	baseURL = "wss://stream.binance.com:9443/ws"
)

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

// WsDepthServe serve websocket depth handler with a symbol
func WsDepthServe(symbol string, handler WsDepthHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@depth", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			// TODO: callback if there is an error
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.UpdateID = j.Get("u").MustInt64()
		bidsLen := len(j.Get("b").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("b").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("a").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("a").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event    string `json:"e"`
	Time     int64  `json:"E"`
	Symbol   string `json:"s"`
	UpdateID int64  `json:"u"`
	Bids     []Bid  `json:"b"`
	Asks     []Ask  `json:"a"`
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(symbol string, interval string, handler WsKlineHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", baseURL, strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}

// WsKline define websocket kline
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsAggTradeHandler handle websocket aggregate trade event
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket aggregate handler with a symbol
func WsAggTradeServe(symbol string, handler WsAggTradeHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", baseURL, strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler)
}

// WsAggTradeEvent define websocket aggregate trade event
type WsAggTradeEvent struct {
	Event                 string `json:"e"`
	Time                  int64  `json:"E"`
	Symbol                string `json:"s"`
	AggTradeID            int64  `json:"a"`
	Price                 string `json:"p"`
	Quantity              string `json:"q"`
	FirstBreakdownTradeID int64  `json:"f"`
	LastBreakdownTradeID  int64  `json:"l"`
	TradeTime             int64  `json:"T"`
	IsBuyerMaker          bool   `json:"m"`
	Placeholder           bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
}

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(listenKey string, handler WsHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s", baseURL, listenKey)
	cfg := newWsConfig(endpoint)
	return wsServe(cfg, handler)
}

type WsUserEventType struct {
	EventType string `json:"e"`
}

type WsUserTradeEvent struct {
	EventType       string          `json:"e"`
	EventTime       int64           `json:"E"`
	Symbol          string          `json:"s"`
	ClientOrderID   string          `json:"c"`
	Side            string          `json:"S"`
	OrderType       string          `json:"o"`
	TimeInForce     string          `json:"f"`
	OrderQuantity   decimal.Decimal `json:"q"`
	OrderPrice      decimal.Decimal `json:"p"`
	StopPrice       decimal.Decimal `json:"P"`
	IcebergQuantity decimal.Decimal `json:"F"`
	//Ignore int
	OriginalClientOrderID    string          `json:"C"`
	ExecutionType            string          `json:"x"`
	OrderStatus              string          `json:"X"`
	RejectReason             string          `json:"r"`
	OrderID                  int64           `json:"i"`
	LastExecutedQuantity     decimal.Decimal `json:"l"`
	CumulativeFilledQuantity decimal.Decimal `json:"z"`
	LastExecutedPrice        decimal.Decimal `json:"L"`
	CommissionAmount         decimal.Decimal `json:"n"`
	CommissionAsset          string          `json:"N"`
	TransactionTime          int64           `json:"T"`
	TradeID                  int64           `json:"t"`
	//Ignore int64
	IsOrderWorking bool `json:"w"`
	IsMaker        bool `json:"m"`
}

type WsUserTradeHandler func(event *WsUserTradeEvent)

func WsUserTradesServe(listenKey string, handler WsUserTradeHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s", baseURL, listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		eventType, eventType_err := jsonparser.GetString(message, "e")
		if eventType_err != nil {
			log.Println("WsUserTradesServe getting eventType error", zap.Error(eventType_err), zap.String("msg", string(message)))
			return
		}

		if eventType == "executionReport" {
			log.Println("WsUserTradesServe raw msg", string(message))
			tradeEvent := new(WsUserTradeEvent)
			decoding_err := jsondecoder_casesensitive.Decode(message, tradeEvent)
			if decoding_err != nil {
				log.Println("WsUserTradesServe decoding json error", zap.Error(decoding_err), zap.String("msg", string(message)))
				return
			}
			handler(tradeEvent)
		} else if eventType == "" {
			log.Println("WsUserTradesServe: Empty EventType!")
		}
	}
	return wsServe(cfg, wsHandler)
}

type WsUserAccountEvent struct {
	EventType            string          `json:"e"`
	EventTime            int64           `json:"E"`
	MakerCommissionRate  float64         `json:"m"`
	TakerCommissionRate  float64         `json:"t"`
	BuyerCommissionRate  float64         `json:"b"`
	SellerCommissionRate float64         `json:"s"`
	CanTrade             bool            `json:"T"`
	CanWithdraw          bool            `json:"W"`
	CanDeposit           bool            `json:"D"`
	LastAccountUpdate    int64           `json:"u"`
	Balances             []BalanceUpdate `json:"B"`
}
type BalanceUpdate struct {
	Asset        string          `json:"a"`
	FreeAmount   decimal.Decimal `json:"f"`
	LockedAmount decimal.Decimal `json:"l"`
}

type WsUserAccountHandler func(event *WsUserAccountEvent)

func WsUserAccountServe(listenKey string, handler WsUserAccountHandler) (chan struct{}, error) {
	endpoint := fmt.Sprintf("%s/%s", baseURL, listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		eventType, eventType_err := jsonparser.GetString(message, "e")
		if eventType_err != nil {
			log.Println("WsUserAccountServe getting eventType error", zap.Error(eventType_err), zap.String("msg", string(message)))
			return
		}

		if eventType == "outboundAccountInfo" {
			accountEvent := new(WsUserAccountEvent)
			decoding_err := jsondecoder_casesensitive.Decode(message, accountEvent)
			if decoding_err != nil {
				log.Println("WsUserAccountServe json decoding error", zap.Error(decoding_err))
				return
			}

			handler(accountEvent)
		} else if eventType == "" {
			log.Println("WsUserAccountServe: Empty EventType!", zap.String("msg", string(message)))
		}
	}
	return wsServe(cfg, wsHandler)
}
