package handler

// User is
type User struct {
	Controller
}

// Create is
// func (uh User) Create(ctx iris.Context) {
// 	u := &models.UserLogin{}
// 	if err := ctx.ReadJSON(u); err != nil {
// 		ctx.StatusCode(iris.StatusBadRequest)
// 		ctx.WriteString(err.Error())
// 		return
// 	}

// 	// u.Password = utils.HashPassword(u.Password)

// 	ResponseJSON(ctx, u)
// }
