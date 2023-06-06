package responses

import "github.com/gin-gonic/gin"

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func SuccessNoData(c *gin.Context, code int) {
	Success(c, code, 0)
}

func Success(c *gin.Context, code int, data any) {
	res := Response{data, ""}
	c.JSON(code, res)
}

func Error(c *gin.Context, code int, err error) {
	res := Response{nil, err.Error()}
	c.JSON(code, res)
}
