package front

import "fmt"

//ErrorFront struct
type ErrorFront struct {
	Status  int      `json:"status"`
	Title   string   `json:"title"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

//ErrorResponse for Front API
type ErrorResponse struct {
	Message ErrorFront `json:"_error"`
}

func (e *ErrorFront) Error() string {
	return fmt.Sprintf("{Status:%d, Title:%s, Message:%s, Detais:%s}",
		e.Status, e.Title, e.Message, e.Details)
}

func (e *ErrorResponse) Error() string {
	return e.Message.Error()
}
