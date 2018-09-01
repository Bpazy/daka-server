package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (d *Date) Scan(src interface{}) error {
	panic("implement me")
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Date) Value() (driver.Value, error) {
	parse, e := time.Parse("2006-01-02", d.String())
	if e != nil {
		panic(e)
	}
	parse = parse.Truncate(time.Hour)
	return parse, nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	result := fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
	return []byte(result), nil
}

func (d *Date) UnmarshalJSON(input []byte) error {
	inputStr := strings.Trim(string(input), "\"")
	dates := strings.Split(inputStr, "-")
	d.Year = parseInt(dates[0])
	d.Month = parseInt(dates[1])
	d.Day = parseInt(dates[2])
	return nil
}

func parseInt(i string) int {
	i2, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}
	return i2
}

type Data struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
	Date     Date   `json:"date"`
}

var db *sql.DB

func init() {
	fmt.Println(os.Args[0])
	db2, err := sql.Open("sqlite3", "./db2.db")
	if err != nil {
		log.Fatal(err)
	}
	db = db2

	rows, err := db2.Query("select COUNT(*) from sqlite_master where name = 'info';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		rows.Scan(&i)
	}
	if i != 0 {
		return
	}

	log.Println("init sqlite and create table...")
	initSqlBytes, err := Asset("res/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	initSql := string(initSqlBytes)
	stmt, err := db.Prepare(initSql)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()
	r.POST("/save", func(c *gin.Context) {
		var data = Data{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Json deserialize failed.",
				"data": err.Error(),
			})
			return
		}
		// TODO github.com/mattn/go-sqlite3 not support (sqlite3 date type)
		stmt, err := db.Prepare("INSERT INTO info(INFO_ID, USERNAME, DISTANCE, DATE) values(?,?,?,?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(uuid.Must(uuid.NewV4()), data.Name, data.Distance, &data.Date)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Insert data failed.",
				"data": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
