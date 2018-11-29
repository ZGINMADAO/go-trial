package controllers

type ApiJson struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func apiResource(status bool, data interface{}, message string) (apiJson *ApiJson) {
	apiJson = &ApiJson{Status: status, Data: data, Message: message}
	return
}
