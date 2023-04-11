package wechatNative

type PayReq struct {
	Mchid       string `json:"mchid"`
	OutTradeNo  string `json:"out_trade_no"`
	Appid       string `json:"appid"`
	Description string `json:"description"`
	NotifyUrl   string `json:"notify_url"`
	Amount      struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
}

type PayRsp struct {
	CodeUrl string `json:"code_url"`
}

func Pay(req *PayReq) (*PayRsp, error) {

}
