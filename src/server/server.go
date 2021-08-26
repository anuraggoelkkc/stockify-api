package server

import (
	"stockify-api/src/support_packs/zerodha"
)

const PORT = ":80"

func Init() {
	r := NewRouter()
	z := zerodha.NewZerodha()
	z.Init()
	r.Run(PORT)
}
