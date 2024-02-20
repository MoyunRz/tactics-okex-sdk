package subaccount

import okex_sdk "tactics-okex-sdk"

type (
	SubAccount struct {
		SubAcct string            `json:"subAcct,omitempty"`
		Label   string            `json:"label,omitempty"`
		Mobile  string            `json:"mobile,omitempty"`
		GAuth   bool              `json:"gAuth"`
		Enable  bool              `json:"enable"`
		TS      okex_sdk.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string            `json:"subAcct,omitempty"`
		Label      string            `json:"label,omitempty"`
		APIKey     string            `json:"apiKey,omitempty"`
		SecretKey  string            `json:"secretKey,omitempty"`
		Passphrase string            `json:"Passphrase,omitempty"`
		Perm       string            `json:"perm,omitempty"`
		IP         string            `json:"ip,omitempty"`
		TS         okex_sdk.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string             `json:"subAcct,omitempty"`
		Ccy     string             `json:"ccy,omitempty"`
		BillID  okex_sdk.JSONInt64 `json:"billId,omitempty"`
		Type    okex_sdk.BillType  `json:"type,omitempty,string"`
		TS      okex_sdk.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransID okex_sdk.JSONInt64 `json:"transId"`
	}
)
