package main

type Foo struct {
	Consign struct {
		Detail          []Detail `json:"detail"`
		TotalAmount     float64  `json:"total_amount"`
		TotalCompensate float64  `json:"total_compensate"`
		TotalFee        float64  `json:"total_fee"`
		TotalManageFee  float64  `json:"total_manage_fee"`
		TotalQuantity   float64  `json:"total_quantity"`
	} `json:"consign"`
	Entrust struct {
		Detail          []Detail `json:"detail"`
		TotalAmount     float64  `json:"total_amount"`
		TotalCompensate float64  `json:"total_compensate"`
		TotalFee        float64  `json:"total_fee"`
		TotalQuantity   float64  `json:"total_quantity"`
	} `json:"entrust"`
	FinishConsign struct {
		Detail             []Detail `json:"detail"`
		TotalQuantity      float64  `json:"total_quantity"`
		TotalTrasferAmount float64  `json:"total_trasfer_amount"`
	} `json:"finish_consign"`
	FinishEntrust struct {
		Detail        []Detail `json:"detail"`
		TotalAmount   float64  `json:"total_amount"`
		TotalQuantity float64  `json:"total_quantity"`
	} `json:"finish_entrust"`
	Unit struct {
		Price         string `json:"price"`
		PriceQuantity string `json:"price_quantity"`
		Quantity      string `json:"quantity"`
	} `json:"unit"`
}

type Detail struct {
	Amount        float64 `json:"amount"`
	Compensate    float64 `json:"compensate"`
	ConsignPrice  float64 `json:"consign_price"`
	DealPrice     float64 `json:"deal_price"`
	Fee           float64 `json:"fee"`
	GoodsTypeName string  `json:"goods_type_name"`
	ManageFee     float64 `json:"manage_fee"`
	Quantity      float64 `json:"quantity"`
	SerialNumber  string  `json:"serial_number"`
	Type          string  `json:"type"`
}
type Detail struct {
	Amount        float64 `json:"amount"`
	Compensate    float64 `json:"compensate"`
	DealPrice     float64 `json:"deal_price"`
	EntrustPrice  float64 `json:"entrust_price"`
	Fee           float64 `json:"fee"`
	GoodsTypeName string  `json:"goods_type_name"`
	Quantity      float64 `json:"quantity"`
	SerialNumber  string  `json:"serial_number"`
	Type          string  `json:"type"`
}
type Detail struct {
	GoodsTypeName string  `json:"goods_type_name"`
	Quantity      float64 `json:"quantity"`
	SerialNumber  string  `json:"serial_number"`
	TradeDate     string  `json:"trade_date"`
	TrasferAmount float64 `json:"trasfer_amount"`
}
type Detail struct {
	Amount        float64 `json:"amount"`
	GoodsTypeName string  `json:"goods_type_name"`
	Quantity      float64 `json:"quantity"`
	SerialNumber  string  `json:"serial_number"`
	TradeDate     string  `json:"trade_date"`
	Type          string  `json:"type"`
}
