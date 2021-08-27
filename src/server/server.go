package server

import "stockify-api/src/support_packs/zerodha"

const PORT = ":8080"

func Init() {
	r := NewRouter()
	r.Run(PORT)
	z := zerodha.NewZerodha()
	z.Init()
}
