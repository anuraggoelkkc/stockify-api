package server

const PORT = ":80"

func Init() {
	r := NewRouter()
	r.Run(PORT)
}
