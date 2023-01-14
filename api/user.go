package api

import (
	"fmt"
	"goober/auth"
	"goober/goober"

	"goober/service/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var service user.RegisterService
	e := c.ShouldBind(&service)
	// e := json.Unmarshal(data, params)
	if e == nil {
		res := service.Register()
		fmt.Println("sh", service)

		c.JSON(200, appendToken(res))
		c.Abort()
		// fmt.Println("er", e)
		return
	} else {
		fmt.Println("should bind er")
		c.JSON(200, goober.WrongResult(e))
	}
}

func appendToken(res *goober.ResponseResult) *goober.ResponseResult {
	if res.Ok {
		m, ok := res.Data.(map[string]any)

		if ok {
			var user = m["user"].(map[string]any)

			token, e := auth.GenToken(user["name"].(string), user["id"].(int64))

			if e == nil {

				m["token"] = token
			}
		}
	}

	return res
}

func Login(c *gin.Context) {
	var service user.LoginService
	e := c.ShouldBind(&service)
	// e := json.Unmarshal(data, params)
	if e == nil {
		res := service.Login()

		c.JSON(200, appendToken(res))
		c.Abort()
		// fmt.Println("er", e)
		return
	} else {
		c.JSON(200, goober.WrongResult(e))
	}

}

func GetUserBaseInfo(c *gin.Context) {
	var service user.GetUserService

	var u, _ = c.Get("user")
	service.Id = int(u.(*auth.JwtUserClaims).Uid)

	c.JSON(200, service.GetBaseInfo())
}

func ModifyUserInfo(c *gin.Context) {
	var service user.ModifyUserInfoService

	e := c.ShouldBind(&service)

	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	var u, _ = c.Get("user")
	service.Id = int(u.(*auth.JwtUserClaims).Uid)

	c.JSON(200, service.Modify())
}
