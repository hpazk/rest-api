package helper

// response
// meta payload

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ResponseFormatter(code int, status string, message string, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
