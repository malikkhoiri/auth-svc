package helper

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type BadResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type M map[string]interface{}
