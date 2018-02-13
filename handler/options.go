package handler

import (
	"anla.io/hound/response"
	"github.com/kataras/iris"
)

// OptionsHandler is
func OptionsHandler(ctx iris.Context) {
	response.JSON(ctx, "hello")
}
