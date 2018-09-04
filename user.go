package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
)

type dakaUser struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func register(username, password string) error {
	smtp, err := db.Prepare("INSERT INTO `daka_user`(`USER_ID`, `USERNAME`) " +
		"SELECT ?, ? FROM DUAL WHERE NOT EXISTS (SELECT NULL FROM daka_user WHERE USERNAME = ?)")
	if err != nil {
		panic(err)
	}
	userId := uuid.Must(uuid.NewV4())
	r, err := smtp.Exec(userId, username, username)
	if err != nil {
		panic(err)
	}
	affectedRows, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	smtp.Close()
	if affectedRows == 0 {
		return errors.New("username exists")
	}
	smtp, err = db.Prepare("insert into daka_user_password (user_id, password) values (?,?)")
	_, err = smtp.Exec(userId, password)
	if err != nil {
		panic(err)
	}
	smtp.Close()
	return nil
}

func findDakaUser(userId string) (d dakaUser, err error) {
	row := db.QueryRow("SELECT USER_ID, USERNAME FROM daka_user where user_id = ?", userId)
	err = row.Scan(&d.UserId, &d.UserName)
	return
}

func findDakaUserByUserName(username string) (*dakaUser, error) {
	row := db.QueryRow("SELECT USER_ID, USERNAME FROM daka_user where USERNAME = ?", username)
	d := dakaUser{}
	err := row.Scan(&d.UserId, &d.UserName)
	if err != nil {
		return nil, err
	}
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
	needLoginFail := func() (statusCode int, r result) {
		statusCode = 419
		r.Msg = "login needed"
		r.Code = "419"
		return
	}
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/api/register" {
			return
		}
		if c.Request.RequestURI == "/api/login" {
			return
		}
		c2, err := c.Cookie(cookieUserId)
		if err != nil && err == http.ErrNoCookie {
			c.AbortWithStatusJSON(needLoginFail())
			return
		}
		cookie := dakaCookie{c2}
		_, err = findDakaUser(cookie.getUserId())
		if err != nil {
			c.AbortWithStatusJSON(needLoginFail())
			return
		}
	}
}
