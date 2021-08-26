package _struct

type Instruments struct {
	InstrumentID string `json:"instrument_token"`
	ExchangeID string `json:"exchange_token"`
	Symbol string `json:"tradingsymbol"`
	Name string `json:"name"`
	Price float64 `json:"last_price"`
	Exchange string `json:"exchange"`
}

type InstrumentDetail struct {
	InstrumentToken   int     `json:"instrument_token"`
	Timestamp         string  `json:"timestamp"`
	LastTradeTime     string  `json:"last_trade_time"`
	LastPrice         float64 `json:"last_price"`
	LastQuantity      int     `json:"last_quantity"`
	BuyQuantity       int     `json:"buy_quantity"`
	SellQuantity      int     `json:"sell_quantity"`
	Volume            int     `json:"volume"`
	AveragePrice      float64 `json:"average_price"`
	Oi                int     `json:"oi"`
	OiDayHigh         int     `json:"oi_day_high"`
	OiDayLow          int     `json:"oi_day_low"`
	NetChange         int     `json:"net_change"`
	LowerCircuitLimit float64 `json:"lower_circuit_limit"`
	UpperCircuitLimit float64 `json:"upper_circuit_limit"`
	Ohlc              struct {
		Open  int     `json:"open"`
		High  float64 `json:"high"`
		Low   float64 `json:"low"`
		Close float64 `json:"close"`
	} `json:"ohlc"`
}

type User struct {
	Id string `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"mail"`
	DeviceID string `json:"device_id"`
}

type Alert struct {
	Id string `json:"alert_id"`
	UserId string `json:"user_id"`
	InstrumentId string `json:"instrument_token"`
	ExchangeId string `json:"exchange_token"`
	Price float64 `json:"target_price"`
	Direction bool `json:"direction"`
}

type Trend struct {}

type Response struct {
	
}