package market

import okex_sdk "github.com/MoyunRz/tactics-okex-sdk"

type (
	GetTickers struct {
		Uly      string                  `json:"uly,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetIndexTickers struct {
		InstID   string `json:"instId,omitempty"`
		QuoteCcy string `json:"quoteCcy,omitempty"`
	}
	GetOrderBook struct {
		InstID string `json:"instId"`
		Sz     int    `json:"sz,omitempty,string"`
	}
	GetCandlesticks struct {
		InstID string           `json:"instId"`
		After  int64            `json:"after,omitempty,string"`
		Before int64            `json:"before,omitempty,string"`
		Limit  int64            `json:"limit,omitempty,string"`
		Bar    okex_sdk.BarSize `json:"bar,omitempty"`
	}
	GetTrades struct {
		InstID string `json:"instId"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
	GetIndexComponents struct {
		Index string `json:"index"`
	}
)
