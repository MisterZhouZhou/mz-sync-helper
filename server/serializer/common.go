package serializer

// 序列化
type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}