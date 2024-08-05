package response

type ResponseStatus struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = 200
	StatusError = 500
)

func Ok() ResponseStatus {
	return ResponseStatus{
		Status: StatusOk,
	}
}

func Error(errMsg string) ResponseStatus {
	return ResponseStatus{
		Status: StatusError,
		Error:  errMsg,
	}
}
