package account

import okex_sdk "github.com/MoyunRz/tactics-okex-sdk"

type (
	Balance struct {
		TotalEq     okex_sdk.JSONFloat64 `json:"totalEq"`
		IsoEq       okex_sdk.JSONFloat64 `json:"isoEq"`
		AdjEq       okex_sdk.JSONFloat64 `json:"adjEq,omitempty"`
		OrdFroz     okex_sdk.JSONFloat64 `json:"ordFroz,omitempty"`
		Imr         okex_sdk.JSONFloat64 `json:"imr,omitempty"`
		Mmr         okex_sdk.JSONFloat64 `json:"mmr,omitempty"`
		MgnRatio    okex_sdk.JSONFloat64 `json:"mgnRatio,omitempty"`
		NotionalUsd okex_sdk.JSONFloat64 `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails    `json:"details,omitempty"`
		UTime       okex_sdk.JSONTime    `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string               `json:"ccy"`
		Eq            okex_sdk.JSONFloat64 `json:"eq"`
		CashBal       okex_sdk.JSONFloat64 `json:"cashBal"`
		IsoEq         okex_sdk.JSONFloat64 `json:"isoEq,omitempty"`
		AvailEq       okex_sdk.JSONFloat64 `json:"availEq,omitempty"`
		DisEq         okex_sdk.JSONFloat64 `json:"disEq"`
		AvailBal      okex_sdk.JSONFloat64 `json:"availBal"`
		FrozenBal     okex_sdk.JSONFloat64 `json:"frozenBal"`
		OrdFrozen     okex_sdk.JSONFloat64 `json:"ordFrozen"`
		Liab          okex_sdk.JSONFloat64 `json:"liab,omitempty"`
		Upl           okex_sdk.JSONFloat64 `json:"upl,omitempty"`
		UplLib        okex_sdk.JSONFloat64 `json:"uplLib,omitempty"`
		CrossLiab     okex_sdk.JSONFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       okex_sdk.JSONFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      okex_sdk.JSONFloat64 `json:"mgnRatio,omitempty"`
		Interest      okex_sdk.JSONFloat64 `json:"interest,omitempty"`
		Twap          okex_sdk.JSONFloat64 `json:"twap,omitempty"`
		MaxLoan       okex_sdk.JSONFloat64 `json:"maxLoan,omitempty"`
		EqUsd         okex_sdk.JSONFloat64 `json:"eqUsd"`
		NotionalLever okex_sdk.JSONFloat64 `json:"notionalLever,omitempty"`
		StgyEq        okex_sdk.JSONFloat64 `json:"stgyEq"`
		IsoUpl        okex_sdk.JSONFloat64 `json:"isoUpl,omitempty"`
		UTime         okex_sdk.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string                  `json:"instId"`
		PosCcy      string                  `json:"posCcy,omitempty"`
		LiabCcy     string                  `json:"liabCcy,omitempty"`
		OptVal      string                  `json:"optVal,omitempty"`
		Ccy         string                  `json:"ccy"`
		PosID       string                  `json:"posId"`
		TradeID     string                  `json:"tradeId"`
		Pos         okex_sdk.JSONFloat64    `json:"pos"`
		AvailPos    okex_sdk.JSONFloat64    `json:"availPos,omitempty"`
		AvgPx       okex_sdk.JSONFloat64    `json:"avgPx"`
		Upl         okex_sdk.JSONFloat64    `json:"upl"`
		UplRatio    okex_sdk.JSONFloat64    `json:"uplRatio"`
		Lever       okex_sdk.JSONFloat64    `json:"lever"`
		LiqPx       okex_sdk.JSONFloat64    `json:"liqPx,omitempty"`
		Imr         okex_sdk.JSONFloat64    `json:"imr,omitempty"`
		Margin      okex_sdk.JSONFloat64    `json:"margin,omitempty"`
		MgnRatio    okex_sdk.JSONFloat64    `json:"mgnRatio"`
		Mmr         okex_sdk.JSONFloat64    `json:"mmr"`
		Liab        okex_sdk.JSONFloat64    `json:"liab,omitempty"`
		Interest    okex_sdk.JSONFloat64    `json:"interest"`
		NotionalUsd okex_sdk.JSONFloat64    `json:"notionalUsd"`
		ADL         okex_sdk.JSONFloat64    `json:"adl"`
		Last        okex_sdk.JSONFloat64    `json:"last"`
		DeltaBS     okex_sdk.JSONFloat64    `json:"deltaBS"`
		DeltaPA     okex_sdk.JSONFloat64    `json:"deltaPA"`
		GammaBS     okex_sdk.JSONFloat64    `json:"gammaBS"`
		GammaPA     okex_sdk.JSONFloat64    `json:"gammaPA"`
		ThetaBS     okex_sdk.JSONFloat64    `json:"thetaBS"`
		ThetaPA     okex_sdk.JSONFloat64    `json:"thetaPA"`
		VegaBS      okex_sdk.JSONFloat64    `json:"vegaBS"`
		VegaPA      okex_sdk.JSONFloat64    `json:"vegaPA"`
		PosSide     okex_sdk.PositionSide   `json:"posSide"`
		MgnMode     okex_sdk.MarginMode     `json:"mgnMode"`
		InstType    okex_sdk.InstrumentType `json:"instType"`
		CTime       okex_sdk.JSONTime       `json:"cTime"`
		UTime       okex_sdk.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex_sdk.EventType `json:"eventType"`
		PTime     okex_sdk.JSONTime  `json:"pTime"`
		UTime     okex_sdk.JSONTime  `json:"uTime"`
		PosData   []*Position        `json:"posData"`
		BalData   []*BalanceDetails  `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okex_sdk.JSONFloat64                 `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex_sdk.JSONTime                    `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string               `json:"ccy"`
		Eq    okex_sdk.JSONFloat64 `json:"eq"`
		DisEq okex_sdk.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string                  `json:"instId"`
		PosCcy      string                  `json:"posCcy,omitempty"`
		Ccy         string                  `json:"ccy"`
		NotionalCcy okex_sdk.JSONFloat64    `json:"notionalCcy"`
		Pos         okex_sdk.JSONFloat64    `json:"pos"`
		NotionalUsd okex_sdk.JSONFloat64    `json:"notionalUsd"`
		PosSide     okex_sdk.PositionSide   `json:"posSide"`
		InstType    okex_sdk.InstrumentType `json:"instType"`
		MgnMode     okex_sdk.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string                  `json:"ccy"`
		InstID    string                  `json:"instId"`
		Notes     string                  `json:"notes"`
		BillID    string                  `json:"billId"`
		OrdID     string                  `json:"ordId"`
		BalChg    okex_sdk.JSONFloat64    `json:"balChg"`
		PosBalChg okex_sdk.JSONFloat64    `json:"posBalChg"`
		Bal       okex_sdk.JSONFloat64    `json:"bal"`
		PosBal    okex_sdk.JSONFloat64    `json:"posBal"`
		Sz        okex_sdk.JSONFloat64    `json:"sz"`
		Pnl       okex_sdk.JSONFloat64    `json:"pnl"`
		Fee       okex_sdk.JSONFloat64    `json:"fee"`
		From      okex_sdk.AccountType    `json:"from,string"`
		To        okex_sdk.AccountType    `json:"to,string"`
		InstType  okex_sdk.InstrumentType `json:"instType"`
		MgnMode   okex_sdk.MarginMode     `json:"MgnMode"`
		Type      okex_sdk.BillType       `json:"type,string"`
		SubType   okex_sdk.BillSubType    `json:"subType,string"`
		TS        okex_sdk.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string                `json:"level"`
		LevelTmp   string                `json:"levelTmp"`
		AcctLv     string                `json:"acctLv"`
		AutoLoan   bool                  `json:"autoLoan"`
		UID        string                `json:"uid"`
		GreeksType okex_sdk.GreekType    `json:"greeksType"`
		PosMode    okex_sdk.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex_sdk.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string                `json:"instId"`
		Lever   okex_sdk.JSONFloat64  `json:"lever"`
		MgnMode okex_sdk.MarginMode   `json:"mgnMode"`
		PosSide okex_sdk.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string               `json:"instId"`
		Ccy     string               `json:"ccy"`
		MaxBuy  okex_sdk.JSONFloat64 `json:"maxBuy"`
		MaxSell okex_sdk.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string               `json:"instId"`
		AvailBuy  okex_sdk.JSONFloat64 `json:"availBuy"`
		AvailSell okex_sdk.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string                `json:"instId"`
		Amt     okex_sdk.JSONFloat64  `json:"amt"`
		PosSide okex_sdk.PositionSide `json:"posSide,string"`
		Type    okex_sdk.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string               `json:"instId"`
		MgnCcy  string               `json:"mgnCcy"`
		Ccy     string               `json:"ccy"`
		MaxLoan okex_sdk.JSONFloat64 `json:"maxLoan"`
		MgnMode okex_sdk.MarginMode  `json:"mgnMode"`
		Side    okex_sdk.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string                  `json:"level"`
		Taker    okex_sdk.JSONFloat64    `json:"taker"`
		Maker    okex_sdk.JSONFloat64    `json:"maker"`
		Delivery okex_sdk.JSONFloat64    `json:"delivery,omitempty"`
		Exercise okex_sdk.JSONFloat64    `json:"exercise,omitempty"`
		Category okex_sdk.FeeCategory    `json:"category,string"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string               `json:"instId"`
		Ccy          string               `json:"ccy"`
		Interest     okex_sdk.JSONFloat64 `json:"interest"`
		InterestRate okex_sdk.JSONFloat64 `json:"interestRate"`
		Liab         okex_sdk.JSONFloat64 `json:"liab"`
		MgnMode      okex_sdk.MarginMode  `json:"mgnMode"`
		TS           okex_sdk.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string               `json:"ccy"`
		InterestRate okex_sdk.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string               `json:"ccy"`
		MaxWd okex_sdk.JSONFloat64 `json:"maxWd"`
	}
)
