package public

import okex_sdk "github.com/MoyunRz/tactics-okex-sdk"

type (
	GetInstruments struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string                  `json:"uly"`
		After    int64                   `json:"after,omitempty,string"`
		Before   int64                   `json:"before,omitempty,string"`
		Limit    int64                   `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		InstID string `json:"instId"`
	}
	GetLimitPrice struct {
		InstID string `json:"instId"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly"`
		Ccy        string  `json:"ccy,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstID   string                  `json:"instId,omitempty"`
		Ccy      string                  `json:"ccy,omitempty"`
		Uly      string                  `json:"uly,omitempty"`
		After    int64                   `json:"after,omitempty,string"`
		Before   int64                   `json:"before,omitempty,string"`
		Limit    int64                   `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		MgnMode  okex_sdk.MarginMode     `json:"mgnMode,omitempty"`
		Alias    okex_sdk.AliasType      `json:"alias,omitempty"`
		State    okex_sdk.OrderState     `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		InstID   string                  `json:"instId,omitempty"`
		Uly      string                  `json:"uly,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		InstID   string                  `json:"instId,omitempty"`
		Uly      string                  `json:"uly,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		TdMode   okex_sdk.TradeMode      `json:"tdMode"`
		Tier     okex_sdk.JSONInt64      `json:"tier,omitempty"`
	}
	GetUnderlying struct {
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	Status struct {
		State string `json:"state,omitempty"`
	}
)
