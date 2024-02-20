package trade

import okex_sdk "tactics-okex-sdk"

type (
	PlaceOrder struct {
		ID         string                `json:"-"`
		InstID     string                `json:"instId"`
		Ccy        string                `json:"ccy,omitempty"`
		ClOrdID    string                `json:"clOrdId,omitempty"`
		Tag        string                `json:"tag,omitempty"`
		ReduceOnly bool                  `json:"reduceOnly,omitempty"`
		Sz         float64               `json:"sz,string"`
		Px         float64               `json:"px,omitempty,string"`
		TdMode     okex_sdk.TradeMode    `json:"tdMode"`
		Side       okex_sdk.OrderSide    `json:"side"`
		PosSide    okex_sdk.PositionSide `json:"posSide,omitempty"`
		OrdType    okex_sdk.OrderType    `json:"ordType"`
		TgtCcy     okex_sdk.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		ID      string `json:"-"`
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID        string  `json:"-"`
		InstID    string  `json:"instId"`
		OrdID     string  `json:"ordId,omitempty"`
		ClOrdID   string  `json:"clOrdId,omitempty"`
		ReqID     string  `json:"reqId,omitempty"`
		NewSz     int64   `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		InstID  string                `json:"instId"`
		Ccy     string                `json:"ccy,omitempty"`
		PosSide okex_sdk.PositionSide `json:"posSide,omitempty"`
		MgnMode okex_sdk.MarginMode   `json:"mgnMode"`
	}
	OrderDetails struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		After    float64                 `json:"after,omitempty,string"`
		Before   float64                 `json:"before,omitempty,string"`
		Begin    string                  `json:"begin,omitempty,string"`
		End      string                  `json:"end,omitempty,string"`
		Limit    float64                 `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
		OrdType  okex_sdk.OrderType      `json:"ordType,omitempty"`
		State    okex_sdk.OrderState     `json:"state,omitempty"`
	}
	TransactionDetails struct {
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		OrdID    string                  `json:"ordId,omitempty"`
		After    float64                 `json:"after,omitempty,string"`
		Before   float64                 `json:"before,omitempty,string"`
		Begin    string                  `json:"begin,omitempty,string"`
		End      string                  `json:"end,omitempty,string"`
		Limit    float64                 `json:"limit,omitempty,string"`
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
	}
	PlaceAlgoOrder struct {
		InstID     string                 `json:"instId"`
		TdMode     okex_sdk.TradeMode     `json:"tdMode"`
		Ccy        string                 `json:"ccy,omitempty"`
		Side       okex_sdk.OrderSide     `json:"side"`
		PosSide    okex_sdk.PositionSide  `json:"posSide,omitempty"`
		OrdType    okex_sdk.AlgoOrderType `json:"ordType"`
		Sz         int64                  `json:"sz,string"`
		ReduceOnly bool                   `json:"reduceOnly,omitempty"`
		TgtCcy     okex_sdk.QuantityType  `json:"tgtCcy,omitempty"`
		StopOrder
		TriggerOrder
		IcebergOrder
		TWAPOrder
	}
	StopOrder struct {
		TpTriggerPx float64 `json:"tpTriggerPx,string,omitempty"`
		TpOrdPx     float64 `json:"tpOrdPx,string,omitempty"`
		SlTriggerPx float64 `json:"slTriggerPx,string,omitempty"`
		SlOrdPx     float64 `json:"slOrdPx,string,omitempty"`
	}
	TriggerOrder struct {
		TriggerPx float64 `json:"triggerPx,string,omitempty"`
		OrdPx     float64 `json:"ordPx,string,omitempty"`
	}
	IcebergOrder struct {
		PxVar    float64 `json:"pxVar,string,omitempty"`
		PxSpread float64 `json:"pxSpread,string,omitempty"`
		SzLimit  int64   `json:"szLimit,string"`
		PxLimit  float64 `json:"pxLimit,string"`
	}
	TWAPOrder struct {
		IcebergOrder
		TimeInterval string `json:"timeInterval"`
	}
	CancelAlgoOrder struct {
		InstID string `json:"instId"`
		AlgoID string `json:"AlgoId"`
	}
	AlgoOrderList struct {
		InstType okex_sdk.InstrumentType `json:"instType,omitempty"`
		Uly      string                  `json:"uly,omitempty"`
		InstID   string                  `json:"instId,omitempty"`
		After    float64                 `json:"after,omitempty,string"`
		Before   float64                 `json:"before,omitempty,string"`
		Limit    float64                 `json:"limit,omitempty,string"`
		OrdType  okex_sdk.AlgoOrderType  `json:"ordType,omitempty"`
		State    okex_sdk.OrderState     `json:"state,omitempty"`
	}
)
