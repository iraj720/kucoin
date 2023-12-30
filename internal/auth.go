package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthInfo struct {
	Code string `json:"code"`
	Data struct {
		Token           string `json:"token"`
		InstanceServers []struct {
			Endpoint     string `json:"endpoint"`
			Encrypt      bool   `json:"encrypt"`
			Protocol     string `json:"protocol"`
			PingInterval int    `json:"pingInterval"`
			PingTimeout  int    `json:"pingTimeout"`
		} `json:"instanceServers"`
	} `json:"data"`
}

func Login() (*AuthInfo, error) {
	request, err := http.NewRequest(http.MethodPost, "https://api.kucoin.com/api/v1/bullet-public", bytes.NewBuffer([]byte(``)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	rsp, err := Request(request, time.Second*10)
	if err != nil {
		return nil, err
	}

	var res AuthInfo
	ress, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(ress, &res)

	return &res, nil
}

type Markets struct {
	Code string   `json:"code"`
	Data []string `json:"data"`
}

func GetMarkets() (*Markets, error) {

	request, err := http.NewRequest(http.MethodGet, "https://api.kucoin.com/api/v1/markets", bytes.NewBuffer([]byte(``)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	rsp, err := Request(request, time.Second*10)
	if err != nil {
		return nil, err
	}

	var res Markets
	ress, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(ress, &res)

	return &res, nil

}

type Symbol struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	FeeCurrency     string `json:"feeCurrency"`
	Market          string `json:"market"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteMinSize    string `json:"quoteMinSize"`
	BaseMaxSize     string `json:"baseMaxSize"`
	QuoteMaxSize    string `json:"quoteMaxSize"`
	BaseIncrement   string `json:"baseIncrement"`
	QuoteIncrement  string `json:"quoteIncrement"`
	PriceIncrement  string `json:"priceIncrement"`
	PriceLimitRate  string `json:"priceLimitRate"`
	MinFunds        string `json:"minFunds"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	EnableTrading   bool   `json:"enableTrading"`
	// map[baseCurrency:DOVI
	// baseIncrement:0.0001
	// baseMaxSize:10000000000
	//baseMinSize:1 enableTrading:true feeCurrency:USDT
	//isMarginEnabled:false market:USDS minFunds:0.1 name:DOVI-USDT
	// priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001
	// quoteMaxSize:99999999 quoteMinSize:0.1 symbol:DOVI-USDT]
}

type Symbols struct {
	Code string   `json:"code"`
	Data []Symbol `json:"data"`
}

var SymBols Symbols

func GetSymbols() (*Symbols, error) {

	request, err := http.NewRequest(http.MethodGet, "https://api.kucoin.com/api/v2/symbols", bytes.NewBuffer([]byte(``)))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	rsp, err := Request(request, time.Second*10)
	if err != nil {
		return nil, err
	}

	res := Symbols{}
	ress, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(ress, &res)
	SymBols = res

	return &res, nil
}

//  map[baseCurrency:DOVI baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:DOVI-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:DOVI-USDT] map[baseCurrency:SEAM baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:0.01 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:SEAM-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:SEAM-USDT] map[baseCurrency:FINC baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:FINC-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:FINC-USDT] map[baseCurrency:VANRY baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:VANRY-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:VANRY-USDT] map[baseCurrency:VANRY baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:BTC isMarginEnabled:false market:BTC minFunds:0.000001 name:VANRY-BTC priceIncrement:0.000000001 priceLimitRate:0.1 quoteCurrency:BTC quoteIncrement:0.000000001 quoteMaxSize:99999999 quoteMinSize:0.000001 symbol:VANRY-BTC] map[baseCurrency:GTT baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:100 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:GTT-USDT priceIncrement:0.000001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.000001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:GTT-USDT] map[baseCurrency:MNDE baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:MNDE-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:MNDE-USDT] map[baseCurrency:COQ baseIncrement:0.0001 baseMaxSize:1000000000000 baseMinSize:100000 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:COQ-USDT priceIncrement:0.000000001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.000000001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:COQ-USDT] map[baseCurrency:IRL baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:IRL-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:IRL-USDT] map[baseCurrency:SOLS baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:0.1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:SOLS-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:SOLS-USDT] map[baseCurrency:POLYX baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:POLYX-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:POLYX-USDT] map[baseCurrency:TAO baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:0.01 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:TAO-USDT priceIncrement:0.01 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.01 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:TAO-USDT] map[baseCurrency:TURT baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:TURT-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:TURT-USDT] map[baseCurrency:BIIS baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:BIIS-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:BIIS-USDT] map[baseCurrency:ARTY baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:ARTY-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:ARTY-USDT] map[baseCurrency:GRAPE baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:GRAPE-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:GRAPE-USDT] map[baseCurrency:MUBI baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:MUBI-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:MUBI-USDT] map[baseCurrency:AA baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:AA-USDT priceIncrement:0.0001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.0001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:AA-USDT] map[baseCurrency:ZOOA baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:ZOOA-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:ZOOA-USDT] map[baseCurrency:ANALOS baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:100 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:ANALOS-USDT priceIncrement:0.000001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.000001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:ANALOS-USDT] map[baseCurrency:MYRO baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:MYRO-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:MYRO-USDT] map[baseCurrency:SILLY baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:1 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:SILLY-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:SILLY-USDT] map[baseCurrency:MOBILE baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:100 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:MOBILE-USDT priceIncrement:0.000001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.000001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:MOBILE-USDT] map[baseCurrency:ROUP baseIncrement:0.0001 baseMaxSize:10000000000 baseMinSize:10 enableTrading:true feeCurrency:USDT isMarginEnabled:false market:USDS minFunds:0.1 name:ROUP-USDT priceIncrement:0.00001 priceLimitRate:0.1 quoteCurrency:USDT quoteIncrement:0.00001 quoteMaxSize:99999999 quoteMinSize:0.1 symbol:ROUP-USDT]

func Request(req *http.Request, timeout time.Duration) (*http.Response, error) {
	tr := http.DefaultTransport
	tc := tr.(*http.Transport).TLSClientConfig
	if tc == nil {
		tc = &tls.Config{InsecureSkipVerify: true}
	} else {
		tc.InsecureSkipVerify = true
	}

	cli := http.DefaultClient
	cli.Transport, cli.Timeout = tr, timeout

	return cli.Do(req)
}
