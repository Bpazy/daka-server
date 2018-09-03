package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type dakaUser struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func findDakaUser(userId string) (d *dakaUser, err error) {
	row := db.QueryRow("SELECT USER_ID, USERNAME FROM daka_user where user_id = ?", userId)
	err = row.Scan(d)
	return
}

func findDakaUserByUserName(username string) (*dakaUser, error) {
	row := db.QueryRow("SELECT USER_ID, USERNAME FROM daka_user where USERNAME = ?", username)
	d := dakaUser{}
	err := row.Scan(&d.UserId, &d.UserName)
	return &d, err
}

func matchPassword(username, password string) bool {
	row := db.QueryRow(`select count(*) from daka_user A join daka_user_password B on A.USER_ID = B.USER_ID 
					where A.username = ? and B.password = ?;`, username, password)
	i := 0
	row.Scan(&i)
	return i != 0
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/api/login" {
			return
		}
		c2, err := c.Cookie(cookieUserId)
		if err != nil && err == http.ErrNoCookie {
			c.JSON(419, "login needed")
			return
		}
		cookie := dakaCookie{c2}
		_, err = findDakaUser(cookie.getUserId())
		if err != nil {
			c.JSON(419, "login needed")
			return
		}
	}
}
