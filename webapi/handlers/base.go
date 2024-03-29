package handlers

import "github.com/gin-gonic/gin"

var Base *BaseHandler

func init() {
	Base = NewBaseHandler()
}

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{

	}
}

type reply struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (b BaseHandler) OK(c *gin.Context, data interface{}) {
	c.JSON(200, reply{200, "success", data})
}

func (b BaseHandler) Fail(c *gin.Context, code int, err error) {
	c.JSON(200, reply{code, err.Error(), nil})
}
