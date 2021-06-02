package domain

import "context"

// PriceUsecase represent the price's usecases
type PriceUsecase interface {
	GetBtcPrice(ctx context.Context, tsyms, fsyms string) (*Price, error)
}

// PriceRepository represent the price's repository contract
type PriceRepository interface {
	Fetch(ctx context.Context, fromVirtualCurrency, toPhysicalCurrency string) ([]CurrencyConversions, error)
	Insert(ctx context.Context, in CurrencyConversions) error
}

// CurrencyConversions is used for fetching results from db
type CurrencyConversions struct {
	ID                   int     `json:"ID"`
	VIRTUALCURRENCYNAME  string  `json:"VIRTUALCURRENCYNAME"`
	PHYSICALCURRENCYNAME string  `json:"PHYSICALCURRENCYNAME"`
	CHANGE24HOUR         float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR      float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR           float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR         float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO       float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR            float64 `json:"LOW24HOUR"`
	HIGH24HOUR           float64 `json:"HIGH24HOUR"`
	PRICE                float64 `json:"PRICE"`
	LASTUPDATE           int64   `json:"LASTUPDATE"`
	SUPPLY               float64 `json:"SUPPLY"`
	MKTCAP               float64 `json:"MKTCAP"`

	STRINGCHANGE24HOUR    string `json:"CHANGE24HOUR"`
	STRINGCHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	STRINGOPEN24HOUR      string `json:"OPEN24HOUR"`
	STRINGVOLUME24HOUR    string `json:"VOLUME24HOUR"`
	STRINGVOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	STRINGLOW24HOUR       string `json:"LOW24HOUR"`
	STRINGHIGH24HOUR      string `json:"HIGH24HOUR"`
	STRINGPRICE           string `json:"PRICE"`
	STRINGLASTUPDATE      string `json:"LASTUPDATE"`
	STRINGSUPPLY          string `json:"SUPPLY"`
	STRINGMKTCAP          string `json:"MKTCAP"`
}

type Raw struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      int64   `json:"LASTUPDATE"`
	SUPPLY          float64 `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}
type Display struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	LOW24HOUR       string `json:"LOW24HOUR"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`
	LASTUPDATE      string `json:"LASTUPDATE"`
	SUPPLY          string `json:"SUPPLY"`
	MKTCAP          string `json:"MKTCAP"`
}
type Price struct {
	RAW     map[string]map[string]Raw     `json:"RAW"`
	DISPLAY map[string]map[string]Display `json:"DISPLAY"`
}
