package funding

import okex_sdk "github.com/MoyunRz/tactics-okex-sdk"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string               `json:"transId"`
		Ccy     string               `json:"ccy"`
		Amt     okex_sdk.JSONFloat64 `json:"amt"`
		From    okex_sdk.AccountType `json:"from,string"`
		To      okex_sdk.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string               `json:"billId"`
		Ccy    string               `json:"ccy"`
		Bal    okex_sdk.JSONFloat64 `json:"bal"`
		BalChg okex_sdk.JSONFloat64 `json:"balChg"`
		Type   okex_sdk.BillType    `json:"type,string"`
		TS     okex_sdk.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string               `json:"addr"`
		Tag      string               `json:"tag,omitempty"`
		Memo     string               `json:"memo,omitempty"`
		PmtID    string               `json:"pmtId,omitempty"`
		Ccy      string               `json:"ccy"`
		Chain    string               `json:"chain"`
		CtAddr   string               `json:"ctAddr"`
		Selected bool                 `json:"selected"`
		To       okex_sdk.AccountType `json:"to,string"`
		TS       okex_sdk.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string                `json:"ccy"`
		Chain string                `json:"chain"`
		TxID  string                `json:"txId"`
		From  string                `json:"from"`
		To    string                `json:"to"`
		DepId string                `json:"depId"`
		Amt   okex_sdk.JSONFloat64  `json:"amt"`
		State okex_sdk.DepositState `json:"state,string"`
		TS    okex_sdk.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		WdID  okex_sdk.JSONInt64   `json:"wdId"`
		Amt   okex_sdk.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string                   `json:"ccy"`
		Chain string                   `json:"chain"`
		TxID  string                   `json:"txId"`
		From  string                   `json:"from"`
		To    string                   `json:"to"`
		Tag   string                   `json:"tag,omitempty"`
		PmtID string                   `json:"pmtId,omitempty"`
		Memo  string                   `json:"memo,omitempty"`
		Amt   okex_sdk.JSONFloat64     `json:"amt"`
		Fee   okex_sdk.JSONFloat64     `json:"fee"`
		WdID  okex_sdk.JSONInt64       `json:"wdId"`
		State okex_sdk.WithdrawalState `json:"state,string"`
		TS    okex_sdk.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string               `json:"ccy"`
		Amt  okex_sdk.JSONFloat64 `json:"amt"`
		Side okex_sdk.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string               `json:"ccy"`
		Amt      okex_sdk.JSONFloat64 `json:"amt"`
		Earnings okex_sdk.JSONFloat64 `json:"earnings"`
	}
)
