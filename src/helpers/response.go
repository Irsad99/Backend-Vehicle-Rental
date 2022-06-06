package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *Response) ResponseJSON(status int, data interface{}) *Response {

	switch status {
	case 200:
		res.Status = status
		res.Message = "OK"
		res.Data = data
	case 201:
		res.Status = status
		res.Message = "Created"
		res.Data = data
	case 204:
		res.Status = status
		res.Message = "No Content"
		res.Data = data
	case 300:
		res.Status = status
		res.Message = "Multiple Choice"
		res.Data = data
	case 304:
		res.Status = status
		res.Message = "Not Modified"
		res.Data = data
	case 400:
		res.Status = status
		res.Message = "Bad Request"
		res.Data = data
	case 401:
		res.Status = status
		res.Message = "Unauthorized"
		res.Data = data
	case 403:
		res.Status = status
		res.Message = "Forbidden"
		res.Data = data
	case 404:
		res.Status = status
		res.Message = "Not Found"
		res.Data = data
	case 500:
		res.Status = status
		res.Message = "Internal Server Error"
		res.Data = data
	case 501:
		res.Status = status
		res.Message = "Bad Gateway"
		res.Data = data
	default:
		res.Status = status
		res.Message = ""
		res.Data = data
	}

	return res
}
