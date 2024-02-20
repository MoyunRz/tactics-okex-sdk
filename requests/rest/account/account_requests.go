package account

import okex_sdk "github.com/MoyunRz/tactics-okex-sdk"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstID   []string                `json:"instId,omitempty"`
		PosID    []string                `json:"posId,omitempty"`
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		Ccy      string                  `json:"ccy,omitempty"`
		After    int64                   `json:"after,omitempty,string"`
		Before   int64                   `json:"before,omitempty,string"`
		Limit    int64                   `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
		MgnMode  okex_sdk.MarginMode     `json:"mgnMode,omitempty"`
		CtType   okex_sdk.ContractType   `json:"ctType,omitempty"`
		Type     okex_sdk.BillType       `json:"type,omitempty,string"`
		SubType  okex_sdk.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PositionMode okex_sdk.PositionType `json:"positionMode"`
	}
	SetLeverage struct {
		Lever   int64                 `json:"lever,string"`
		InstID  string                `json:"instId,omitempty"`
		Ccy     string                `json:"ccy,omitempty"`
		MgnMode okex_sdk.MarginMode   `json:"mgnMode"`
		PosSide okex_sdk.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string             `json:"ccy,omitempty"`
		Px     float64            `json:"px,string,omitempty"`
		InstID []string           `json:"instId"`
		TdMode okex_sdk.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string             `json:"ccy,omitempty"`
		InstID     string             `json:"instId"`
		ReduceOnly bool               `json:"reduceOnly,omitempty"`
		TdMode     okex_sdk.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		InstID     string                `json:"instId"`
		Amt        float64               `json:"amt,string"`
		PosSide    okex_sdk.PositionSide `json:"posSide"`
		ActionType okex_sdk.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		InstID  []string            `json:"instId"`
		MgnMode okex_sdk.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		InstID  string              `json:"instId"`
		MgnCcy  string              `json:"mgnCcy,omitempty"`
		MgnMode okex_sdk.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		InstID   string                  `json:"instId,omitempty"`
		Uly      string                  `json:"uly,omitempty"`
		Category okex_sdk.FeeCategory    `json:"category,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		InstID  string              `json:"instId,omitempty"`
		Ccy     string              `json:"ccy,omitempty"`
		After   int64               `json:"after,omitempty,string"`
		Before  int64               `json:"before,omitempty,string"`
		Limit   int64               `json:"limit,omitempty,string"`
		MgnMode okex_sdk.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		GreeksType okex_sdk.GreekType `json:"greeksType"`
	}
)
