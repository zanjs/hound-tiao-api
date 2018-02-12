package handler

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// UserHand is
type UserHand struct{}

// MyCustomClaims is
type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

var mySigningKey = []byte("MySecret22")

// Create is
func (u UserHand) Create(ctx iris.Context) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2018, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// Create the Claims
	claims := MyCustomClaims{
		"bar2",
		jwt.StandardClaims{
			ExpiresAt: 999990,
			Issuer:    "user2",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	fmt.Println(err)

	ResponseJSON(ctx, ss)
}
