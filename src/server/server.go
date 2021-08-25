package server

import (
	"stockify-api/src/zerodha"
)

const PORT = ":80"

func Init() {
	r := NewRouter()
	zerodha.Init()
	r.Run(PORT)
}
