package handler

import (
	"anla.io/hound/models"
	"anla.io/hound/response"
	"anla.io/hound/utils"
	"github.com/houndgo/suuid"
	"github.com/kataras/iris"
)

type (
	// Register 注册接口
	Register struct{}
)

// Add is add user
func (re Register) Add(ctx iris.Context) {
	u := &models.UserLogin{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Username == "" {
		response.JSONError(ctx, "Username where?")
		return
	}

	if u.Password == "" {
		response.JSONError(ctx, "Password where?")
		return
	}

	user, _ := models.User{}.GetByUsername(u.Username)

	if user.ID != 0 {
		response.JSONError(ctx, "用户名存在")
		return
	}

	user.Username = u.Username
	user.Password = utils.HashPassword(u.Password)
	user.UID = suuid.New().String()

	err := models.User{}.Create(&user)

	if err != nil {
		response.JSONError(ctx, "注册失败")
	}

	response.JSON(ctx, user)
}
