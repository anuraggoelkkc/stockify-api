package _struct

type Instruments struct {
	Instrument_ID string `json:"instrument_token"`
	Exchange_ID string `json:"exchange_token"`
	Symbol string `json:"tradingsymbol"`
	Name string `json:"name"`
	Price float64 `json:"last_price"`
	Exchange string `json:"exchange"`
}

type User struct {
	ID string `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"mail"`
	DeviceID string `json:"device_id"`
}

type Alert struct {
	ID string `json:"alert_id"`
	User_ID string `json:"user_id"`
	Instrument_ID string `json:"instrument_token"`
	Exchange_ID string `json:"exchange_token"`
	Price float64 `json:"target_price"`
	Direction bool `json:"direction"`
}

type Trend struct {}