package handler

import (
	"anla.io/hound/response"
	"github.com/kataras/iris"
)

type (
	Article struct{}
)

func (a Article) Create(ctx iris.Context) {
	response.JSON(ctx, "s")
}
