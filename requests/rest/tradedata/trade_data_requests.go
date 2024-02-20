package tradedata

import "chain-tactics/okex-sdk"

type (
	GetTakerVolume struct {
		Ccy      string                  `json:"ccy"`
		Begin    int64                   `json:"before,omitempty,string"`
		End      int64                   `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		Period   okex_sdk.BarSize        `json:"period,string,omitempty"`
	}
	GetRatio struct {
		Ccy    string           `json:"ccy"`
		Begin  int64            `json:"before,omitempty,string"`
		End    int64            `json:"limit,omitempty,string"`
		Period okex_sdk.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		Ccy     string           `json:"ccy"`
		ExpTime string           `json:"expTime"`
		Period  okex_sdk.BarSize `json:"period,string,omitempty"`
	}
)
