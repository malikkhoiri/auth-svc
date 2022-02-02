package helper

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BadResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type M map[string]interface{}
