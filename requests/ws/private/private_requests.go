package private

import "chain-tactics/okex-sdk"

type (
	Account struct {
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	Order struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
)
