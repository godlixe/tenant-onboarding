package httpx

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
