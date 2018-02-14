package handler

import (
	"anla.io/hound/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// Controller is base controller
type Controller struct{}

// GetUser 获取用户信息
func (ctl Controller) GetUser(ctx iris.Context) models.User {
	user := models.User{}
	userJwt := ctx.Values().Get("jwt").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	user.ID = userID
	user.Username = claims["username"].(string)
	return user
}
