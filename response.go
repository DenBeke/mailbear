package mailbear

type mailbearResponse struct {
	Message string `json:"message"`
}

func mailbearRespone(message string) mailbearResponse {
	return mailbearResponse{Message: message}
}
