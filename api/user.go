package api

import (
	"fmt"
	"goblog/auth"
	"goblog/rep"
	"goblog/service/user"

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
		c.JSON(200, rep.BuildFatalResponse(e))
	}
}

func appendToken(res *rep.Response) *rep.Response {
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
		c.JSON(200, rep.BuildFatalResponse(e))
	}

}