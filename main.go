package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strings"
)

var db *sqlx.DB
var port *string

const (
	cookieUserId = "cookieUserId"
	OK           = "OK"
	FAILED       = "FAILED"
)

func init() {
	sqlUserName := flag.String("sqlUserName", "", "mysql user name")
	sqlPassword := flag.String("sqlPassword", "", "mysql user password")
	sqlUrl := flag.String("sqlUrl", "", "mysql url")
	sqlDatabase := flag.String("sqlDatabase", "", "database")
	port = flag.String("port", ":8080", "serve port")
	flag.Parse()

	db2, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *sqlUserName, *sqlPassword, *sqlUrl, *sqlDatabase))
	if err != nil {
		log.Fatal("connect database error", err)
	}
	db = db2

	rows, err := db2.Query("SELECT count(*) FROM information_schema.TABLES WHERE table_name = 'daka_info';")
	if err != nil {
		log.Fatal("query table exists error", err)
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		rows.Scan(&i)
	}
	if i != 0 {
		return
	}

	log.Println("init mysql and create table...")
	initSqlBytes, err := Asset("res/init.sql")
	if err != nil {
		log.Fatal("load init.sql error", err)
	}
	initSqls := strings.Split(strings.TrimSpace(string(initSqlBytes)), ";")
	for _, initSql := range initSqls {
		initSql := strings.TrimSpace(initSql)
		if len(initSql) == 0 {
			continue
		}
		stmt, err := db.Prepare(initSql)
		if err != nil {
			log.Fatal("prepare init sql error", err)
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Fatal("exec init sql error", err)
		}
		stmt.Close()
	}
}

func main() {
	r := gin.Default()
	r.Use(userMiddleware())

	api := r.Group("/api")
	{
		api.POST("/save", saveHandler())
		api.POST("/list", listHandler())
		api.POST("/login", loginHandler())
		api.POST("/register", registerHandler())
	}
	r.Run(*port)
}

type result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ok(msg string, data interface{}) (httpStatus int, r result) {
	httpStatus = http.StatusOK
	r.Msg = msg
	r.Data = data
	r.Code = OK
	return
}

func fail(msg string, data interface{}) (httpStatus int, r result) {
	httpStatus = http.StatusOK
	r.Msg = msg
	r.Data = data
	r.Code = FAILED
	return
}

func registerHandler() gin.HandlerFunc {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		rr := RegisterRequest{}
		if err := c.BindJSON(&rr); err != nil {
			c.JSON(fail("Json deserialize failed", err.Error()))
			return
		}
		if err := register(rr.Username, rr.Password); err != nil {
			c.JSON(fail(err.Error(), ""))
			return
		}
		c.JSON(ok("register success", ""))
	}
}

func loginHandler() gin.HandlerFunc {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		lr := LoginRequest{}
		if err := c.BindJSON(&lr); err != nil {
			c.JSON(fail("Json deserialize failed", err.Error()))
			return
		}

		user, err := findDakaUserByUserName(lr.Username)
		if err != nil {
			c.JSON(fail("user name or password is incorrect", err.Error()))
			return
		}
		if !matchPassword(lr.Username, lr.Password) {
			c.JSON(fail("user name or password is incorrect", nil))
			return
		}
		c.SetCookie(cookieUserId, user.UserId, 60*60*8, "/", "", false, false)
		c.JSON(ok("login success", ""))
	}
}

type Data struct {
	InfoId   string `json:"info_id" db:"info_id"`
	Name     string `json:"name"`
	Distance int    `json:"distance"`
	Date     Date   `json:"date"`
}

func saveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = Data{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(fail("Json deserialize failed.", err.Error()))
			return
		}
		data.InfoId = uuid.Must(uuid.NewV4()).String()
		nstmt, err := db.PrepareNamed("INSERT INTO daka_info(INFO_ID, NAME, DISTANCE, DATE) values(:info_id,:name,:distance,:date)")
		if err != nil {
			panic(err)
		}
		defer nstmt.Close()

		_, err = nstmt.Exec(&data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}

type PaginationRequest struct {
	Start int `json:"start"`
	Size  int `json:"size"`
}

type PaginationResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func listHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := PaginationRequest{}
		if err := c.BindJSON(&pagination); err != nil {
			c.JSON(fail("Json deserialize failed.", ""))
			return
		}

		total := 0
		db.Get(&total, "SELECT COUNT(*) FROM daka_info;")

		rows, err := db.NamedQuery("SELECT info_id, name, distance, date FROM daka_info ORDER BY create_time DESC LIMIT :start,:size", &pagination)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		dataList := make([]Data, 0)
		if err = sqlx.StructScan(rows, &dataList); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, PaginationResponse{Total: total, Data: dataList})
	}
}
