package trade

import (
	"chain-tactics/okex-sdk"
)

type (
	PlaceOrder struct {
		ClOrdID string               `json:"clOrdId"`
		Tag     string               `json:"tag"`
		SMsg    string               `json:"sMsg"`
		SCode   okex_sdk.JSONInt64   `json:"sCode"`
		OrdID   okex_sdk.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string               `json:"ordId"`
		ClOrdID string               `json:"clOrdId"`
		SMsg    string               `json:"sMsg"`
		SCode   okex_sdk.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string               `json:"ordId"`
		ClOrdID string               `json:"clOrdId"`
		ReqID   string               `json:"reqId"`
		SMsg    string               `json:"sMsg"`
		SCode   okex_sdk.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string                `json:"instId"`
		PosSide okex_sdk.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string                  `json:"instId"`
		Ccy         string                  `json:"ccy"`
		OrdID       string                  `json:"ordId"`
		ClOrdID     string                  `json:"clOrdId"`
		TradeID     string                  `json:"tradeId"`
		Tag         string                  `json:"tag"`
		Category    string                  `json:"category"`
		FeeCcy      string                  `json:"feeCcy"`
		RebateCcy   string                  `json:"rebateCcy"`
		Px          okex_sdk.JSONFloat64    `json:"px"`
		Sz          okex_sdk.JSONFloat64    `json:"sz"`
		Pnl         okex_sdk.JSONFloat64    `json:"pnl"`
		AccFillSz   okex_sdk.JSONFloat64    `json:"accFillSz"`
		FillPx      okex_sdk.JSONFloat64    `json:"fillPx"`
		FillSz      okex_sdk.JSONFloat64    `json:"fillSz"`
		FillTime    okex_sdk.JSONFloat64    `json:"fillTime"`
		AvgPx       okex_sdk.JSONFloat64    `json:"avgPx"`
		Lever       okex_sdk.JSONFloat64    `json:"lever"`
		TpTriggerPx okex_sdk.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx     okex_sdk.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx okex_sdk.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx     okex_sdk.JSONFloat64    `json:"slOrdPx"`
		Fee         okex_sdk.JSONFloat64    `json:"fee"`
		Rebate      okex_sdk.JSONFloat64    `json:"rebate"`
		State       okex_sdk.OrderState     `json:"state"`
		TdMode      okex_sdk.TradeMode      `json:"tdMode"`
		PosSide     okex_sdk.PositionSide   `json:"posSide"`
		Side        okex_sdk.OrderSide      `json:"side"`
		OrdType     okex_sdk.OrderType      `json:"ordType"`
		InstType    okex_sdk.InstrumentType `json:"instType"`
		TgtCcy      okex_sdk.QuantityType   `json:"tgtCcy"`
		UTime       okex_sdk.JSONTime       `json:"uTime"`
		CTime       okex_sdk.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string                  `json:"instId"`
		OrdID    string                  `json:"ordId"`
		TradeID  string                  `json:"tradeId"`
		ClOrdID  string                  `json:"clOrdId"`
		BillID   string                  `json:"billId"`
		Tag      okex_sdk.JSONFloat64    `json:"tag"`
		FillPx   okex_sdk.JSONFloat64    `json:"fillPx"`
		FillSz   okex_sdk.JSONFloat64    `json:"fillSz"`
		FeeCcy   okex_sdk.JSONFloat64    `json:"feeCcy"`
		Fee      okex_sdk.JSONFloat64    `json:"fee"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		Side     okex_sdk.OrderSide      `json:"side"`
		PosSide  okex_sdk.PositionSide   `json:"posSide"`
		ExecType okex_sdk.OrderFlowType  `json:"execType"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string             `json:"algoId"`
		SMsg   string             `json:"sMsg"`
		SCode  okex_sdk.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string             `json:"algoId"`
		SMsg   string             `json:"sMsg"`
		SCode  okex_sdk.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string                  `json:"instId"`
		Ccy          string                  `json:"ccy"`
		OrdID        string                  `json:"ordId"`
		AlgoID       string                  `json:"algoId"`
		ClOrdID      string                  `json:"clOrdId"`
		TradeID      string                  `json:"tradeId"`
		Tag          string                  `json:"tag"`
		Category     string                  `json:"category"`
		FeeCcy       string                  `json:"feeCcy"`
		RebateCcy    string                  `json:"rebateCcy"`
		TimeInterval string                  `json:"timeInterval"`
		Px           okex_sdk.JSONFloat64    `json:"px"`
		PxVar        okex_sdk.JSONFloat64    `json:"pxVar"`
		PxSpread     okex_sdk.JSONFloat64    `json:"pxSpread"`
		PxLimit      okex_sdk.JSONFloat64    `json:"pxLimit"`
		Sz           okex_sdk.JSONFloat64    `json:"sz"`
		SzLimit      okex_sdk.JSONFloat64    `json:"szLimit"`
		ActualSz     okex_sdk.JSONFloat64    `json:"actualSz"`
		ActualPx     okex_sdk.JSONFloat64    `json:"actualPx"`
		Pnl          okex_sdk.JSONFloat64    `json:"pnl"`
		AccFillSz    okex_sdk.JSONFloat64    `json:"accFillSz"`
		FillPx       okex_sdk.JSONFloat64    `json:"fillPx"`
		FillSz       okex_sdk.JSONFloat64    `json:"fillSz"`
		FillTime     okex_sdk.JSONFloat64    `json:"fillTime"`
		AvgPx        okex_sdk.JSONFloat64    `json:"avgPx"`
		Lever        okex_sdk.JSONFloat64    `json:"lever"`
		TpTriggerPx  okex_sdk.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx      okex_sdk.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx  okex_sdk.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx      okex_sdk.JSONFloat64    `json:"slOrdPx"`
		OrdPx        okex_sdk.JSONFloat64    `json:"ordPx"`
		Fee          okex_sdk.JSONFloat64    `json:"fee"`
		Rebate       okex_sdk.JSONFloat64    `json:"rebate"`
		State        okex_sdk.OrderState     `json:"state"`
		TdMode       okex_sdk.TradeMode      `json:"tdMode"`
		ActualSide   okex_sdk.PositionSide   `json:"actualSide"`
		PosSide      okex_sdk.PositionSide   `json:"posSide"`
		Side         okex_sdk.OrderSide      `json:"side"`
		OrdType      okex_sdk.AlgoOrderType  `json:"ordType"`
		InstType     okex_sdk.InstrumentType `json:"instType"`
		TgtCcy       okex_sdk.QuantityType   `json:"tgtCcy"`
		CTime        okex_sdk.JSONTime       `json:"cTime"`
		TriggerTime  okex_sdk.JSONTime       `json:"triggerTime"`
	}
)
