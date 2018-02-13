package response

import (
	"github.com/kataras/iris"
)

// Response : JSON Response Object
type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// JSON ...
func JSON(ctx iris.Context, d interface{}) {
	ctx.JSON(&Response{
		Success: true,
		Data:    d,
	})
	ctx.Application().Logger().Info("response is success data: " + d.(string))
}

// JSONError ...
func JSONError(ctx iris.Context, err string) {
	ctx.JSON(&Response{
		Success: false,
		Message: err,
	})
	ctx.Application().Logger().Info("response is error : " + err)
}

// JSONBad is ...
func JSONBad(ctx iris.Context, err string) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(&Response{
		Success: false,
		Message: err,
	})
	ctx.Application().Logger().Info("response is bad data: " + err)
}
