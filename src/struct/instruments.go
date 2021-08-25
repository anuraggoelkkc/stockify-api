package _struct

type Instruments struct {
	Instrument_ID string `json:"instrument_token"`
	Exchange_ID string `json:"exchange_token"`
	Symbol string `json:"tradingsymbol"`
	Name string `json:"name"`
	Price float64 `json:"last_price"`
	Exchange string `json:"exchange"`
}
