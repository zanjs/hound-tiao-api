package handler

import (
	"github.com/kataras/iris"
)

// OptionsHandler is
func OptionsHandler(ctx iris.Context) {
	ResponseJSON(ctx, "hello")
}
