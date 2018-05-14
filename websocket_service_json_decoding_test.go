package binance

import "testing"
import (
	"log"

	"github.com/json-iterator/go"
)

func TestJsonDecoding(t *testing.T) {
	var msg = []byte(`{"e":"executionReport","E":1524984672572,"s":"BTCUSDT","c":"web_48793d9220be4fde9d05ad36320ada3a","S":"SELL","o":"LIMIT","f":"GTC","q":"0.00684500","p":"20000.00000000","P":"0.00000000","F":"0.00000000","g":-1,"C":"web_81b25bddeaeb4563b2e69de939b152e2","x":"CANCELED","X":"CANCELED","r":"NONE","i":94216561,"l":"0.00000000","z":"0.00000000","L":"0.00000000","n":"0","N":null,"T":1524984672568,"t":-1,"I":227596183,"w":false,"m":false,"M":false,"O":-1,"Z":"-0.00000001"}`)
	var userEvent WsUserTradeEvent
	//if decoding_err := json.Unmarshal(msg, &userEvent); decoding_err != nil {
	//	t.Error(decoding_err)
	//}
	//extra.SetNamingStrategy(func(s string) string {
	//	log.Println("s", s)
	//	return s
	//})
	jsonConfig := jsoniter.Config{
		OnlyTaggedField: true,
	}.Froze()
	if decoding_err := jsonConfig.Unmarshal(msg, &userEvent); decoding_err != nil {
		t.Error(decoding_err)
	}

	log.Println(userEvent)
	log.Println(userEvent.OrderType)

	if userEvent.EventType != "executionReport" {
		t.Error("error when parsing")
	}
	if userEvent.Symbol != "BTCUSDT" {
		t.Error("Error when parsing")
	}

}
