package publicdata
import okex_sdk "tactics-okex-sdk"
type (
	Instrument struct {
		InstID    string                   `json:"instId"`
		Uly       string                   `json:"uly,omitempty"`
		BaseCcy   string                   `json:"baseCcy,omitempty"`
		QuoteCcy  string                   `json:"quoteCcy,omitempty"`
		SettleCcy string                   `json:"settleCcy,omitempty"`
		CtValCcy  string                   `json:"ctValCcy,omitempty"`
		CtVal     okex_sdk.JSONFloat64     `json:"ctVal,omitempty"`
		CtMult    okex_sdk.JSONFloat64     `json:"ctMult,omitempty"`
		Stk       okex_sdk.JSONFloat64     `json:"stk,omitempty"`
		TickSz    okex_sdk.JSONFloat64     `json:"tickSz,omitempty"`
		LotSz     okex_sdk.JSONFloat64     `json:"lotSz,omitempty"`
		MinSz     okex_sdk.JSONFloat64     `json:"minSz,omitempty"`
		Lever     okex_sdk.JSONFloat64     `json:"lever"`
		InstType  okex_sdk.InstrumentType  `json:"instType"`
		Category  okex_sdk.FeeCategory     `json:"category,string"`
		OptType   okex_sdk.OptionType      `json:"optType,omitempty"`
		ListTime  okex_sdk.JSONTime        `json:"listTime"`
		ExpTime   okex_sdk.JSONTime        `json:"expTime,omitempty"`
		CtType    okex_sdk.ContractType    `json:"ctType,omitempty"`
		Alias     okex_sdk.AliasType       `json:"alias,omitempty"`
		State     okex_sdk.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex_sdk.JSONTime                 `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                        `json:"instId"`
		Px     okex_sdk.JSONFloat64          `json:"px"`
		Type   okex_sdk.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string                  `json:"instId"`
		Oi       okex_sdk.JSONFloat64    `json:"oi"`
		OiCcy    okex_sdk.JSONFloat64    `json:"oiCcy"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string                  `json:"instId"`
		InstType        okex_sdk.InstrumentType `json:"instType"`
		FundingRate     okex_sdk.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okex_sdk.JSONFloat64    `json:"NextFundingRate"`
		FundingTime     okex_sdk.JSONTime       `json:"fundingTime"`
		NextFundingTime okex_sdk.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string                  `json:"instId"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		BuyLmt   okex_sdk.JSONFloat64    `json:"buyLmt"`
		SellLmt  okex_sdk.JSONFloat64    `json:"sellLmt"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string                  `json:"instId"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		SettlePx okex_sdk.JSONFloat64    `json:"settlePx"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string                  `json:"instId"`
		Uly      string                  `json:"uly"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		Delta    okex_sdk.JSONFloat64    `json:"delta"`
		Gamma    okex_sdk.JSONFloat64    `json:"gamma"`
		Vega     okex_sdk.JSONFloat64    `json:"vega"`
		Theta    okex_sdk.JSONFloat64    `json:"theta"`
		DeltaBS  okex_sdk.JSONFloat64    `json:"deltaBS"`
		GammaBS  okex_sdk.JSONFloat64    `json:"gammaBS"`
		VegaBS   okex_sdk.JSONFloat64    `json:"vegaBS"`
		ThetaBS  okex_sdk.JSONFloat64    `json:"thetaBS"`
		Lever    okex_sdk.JSONFloat64    `json:"lever"`
		MarkVol  okex_sdk.JSONFloat64    `json:"markVol"`
		BidVol   okex_sdk.JSONFloat64    `json:"bidVol"`
		AskVol   okex_sdk.JSONFloat64    `json:"askVol"`
		RealVol  okex_sdk.JSONFloat64    `json:"realVol"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string               `json:"ccy"`
		Amt          okex_sdk.JSONFloat64 `json:"amt"`
		DiscountLv   okex_sdk.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo      `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex_sdk.JSONInt64 `json:"discountRate"`
		MaxAmt       okex_sdk.JSONInt64 `json:"maxAmt"`
		MinAmt       okex_sdk.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okex_sdk.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex_sdk.InstrumentType   `json:"instType"`
		TotalLoss okex_sdk.JSONFloat64      `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string                `json:"ccy,omitempty"`
		Side    okex_sdk.OrderSide    `json:"side"`
		OosSide okex_sdk.PositionSide `json:"posSide"`
		BkPx    okex_sdk.JSONFloat64  `json:"bkPx"`
		Sz      okex_sdk.JSONFloat64  `json:"sz"`
		BkLoss  okex_sdk.JSONFloat64  `json:"bkLoss"`
		TS      okex_sdk.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string                  `json:"instId"`
		InstType okex_sdk.InstrumentType `json:"instType"`
		MarkPx   okex_sdk.JSONFloat64    `json:"markPx"`
		TS       okex_sdk.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string                  `json:"instId"`
		Uly          string                  `json:"uly,omitempty"`
		InstType     okex_sdk.InstrumentType `json:"instType"`
		Tier         okex_sdk.JSONInt64      `json:"tier"`
		MinSz        okex_sdk.JSONFloat64    `json:"minSz"`
		MaxSz        okex_sdk.JSONFloat64    `json:"maxSz"`
		Mmr          okex_sdk.JSONFloat64    `json:"mmr"`
		Imr          okex_sdk.JSONFloat64    `json:"imr"`
		OptMgnFactor okex_sdk.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex_sdk.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex_sdk.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okex_sdk.JSONFloat64    `json:"maxLever"`
		TS           okex_sdk.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string               `json:"ccy"`
		Rate  okex_sdk.JSONFloat64 `json:"rate"`
		Quota okex_sdk.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string               `json:"level"`
		IrDiscount    okex_sdk.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int                  `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string            `json:"title"`
		State       string            `json:"state"`
		Href        string            `json:"href"`
		ServiceType string            `json:"serviceType"`
		System      string            `json:"system"`
		ScheDesc    string            `json:"scheDesc"`
		Begin       okex_sdk.JSONTime `json:"begin"`
		End         okex_sdk.JSONTime `json:"end"`
	}
)
