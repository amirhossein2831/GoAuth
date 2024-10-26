package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (r Response) Send(c *gin.Context) {
	c.JSON(r.Status, r)
}
