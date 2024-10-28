package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	Send()
}

type Response struct {
	c            *gin.Context
	IsSuccessful bool           `json:"is_successful"`
	StatusCode   int            `json:"status_code"`
	Error        string         `json:"error"`
	Data         map[string]any `json:"data,omitempty"`
}

// NewResponse initializes a new response builder with default values
func NewResponse(c *gin.Context) *Response {
	return &Response{
		IsSuccessful: false,
		StatusCode:   http.StatusBadRequest,
		c:            c,
	}
}

// SetStatusCode sets the status code of the response
func (response *Response) SetStatusCode(statusCode int) *Response {
	response.StatusCode = statusCode
	return response
}

// SetError sets the message of the response
func (response *Response) SetError(err error) *Response {
	response.Error = err.Error()
	return response
}

// SetData sets the data of the response
func (response *Response) SetData(data map[string]any) *Response {
	response.Data = data
	return response
}

// Send sends the constructed response to the client
func (response *Response) Send() {
	response.IsSuccessful = response.StatusCode >= 200 && response.StatusCode < 300
	response.c.JSON(response.StatusCode, response)
}
