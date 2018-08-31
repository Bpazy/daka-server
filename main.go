package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
	"google.golang.org/genproto/googleapis/type/date"
	"log"
	"net/http"
)

type Data struct {
	Name     string    `json:"name"`
	Distance int       `json:"distance"`
	Date     date.Date `json:"date" date_format:""`
}

var db *sql.DB

func init() {
	db2, err := sql.Open("sqlite3", "./db2.db")
	if err != nil {
		log.Fatal(err)
	}
	db = db2
}

func main() {
	r := gin.Default()
	r.POST("/save", func(c *gin.Context) {
		var data = Data{}
		if err := c.BindJSON(&data); err != nil {
			handleErr(c, err)
			return
		}
		stmt, err := db.Prepare("INSERT INTO info(INFO_ID, USERNAME, DISTANCE, DATE) values(?,?,?,?)")
		if err != nil {
			panic(err)
		}
		stmt.Exec(uuid.Must(uuid.NewV4()), data.Name, data.Distance, data.Date)
		c.JSON(http.StatusOK, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func json() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func handleErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Json deserialize failed.",
		"data": err.Error(),
	})
}
