package handler

import (
	"github.com/kataras/iris"
)

// IndexHand is
func IndexHand(ctx iris.Context) {
	ResponseJSON(ctx, "hello boy!")
}
