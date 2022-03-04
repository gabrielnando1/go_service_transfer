package dto_common_response

type GenericResponse struct {
	ErrorMessage *string     `json:"error_message"`
	Success      bool        `json:"success"`
	Content      interface{} `json:"content"`
}
