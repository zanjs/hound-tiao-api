package handler

import "github.com/kataras/iris"

// JWTError is
func JWTError(ctx iris.Context, str string) {
	ResponseJSONError(ctx, str)
}
