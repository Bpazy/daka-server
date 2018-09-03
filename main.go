package main

import (
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
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
	dateStr := string(src.([]uint8))
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Year = t.Year()
	d.Month = int(t.Month())
	d.Day = t.Day()
	return nil
}

func (d Date) Value() (driver.Value, error) {
	parse, err := time.Parse("2006-01-02", d.String())
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Date) MarshalJSON() ([]byte, error) {
	result := fmt.Sprintf("\"%d-%02d-%02d\"", d.Year, d.Month, d.Day)
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
	InfoId   string `json:"info_id"`
	Name     string `json:"name"`
	Distance int    `json:"distance"`
	Date     Date   `json:"date"`
}

var db *sql.DB
var port *string

func init() {
	sqlUserName := flag.String("sqlUserName", "", "mysql user name")
	sqlPassword := flag.String("sqlPassword", "", "mysql user password")
	sqlUrl := flag.String("sqlUrl", "", "mysql url")
	sqlDatabase := flag.String("sqlDatabase", "", "database")
	port = flag.String("port", ":8080", "serve port")
	flag.Parse()

	db2, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *sqlUserName, *sqlPassword, *sqlUrl, *sqlDatabase))
	if err != nil {
		log.Fatal(err)
	}
	db = db2

	rows, err := db2.Query("SELECT count(*) FROM information_schema.TABLES WHERE table_name = 'info';")
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

	log.Println("init mysql and create table...")
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

	api := r.Group("/api")
	{
		api.POST("/save", saveHandler())
		api.POST("/list", listHandler())
	}
	r.Run(*port) // listen and serve on 0.0.0.0:8080
}

func saveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = Data{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Json deserialize failed.",
				"data": err.Error(),
			})
			return
		}
		stmt, err := db.Prepare("INSERT INTO info(INFO_ID, NAME, DISTANCE, DATE) values(?,?,?,?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(uuid.Must(uuid.NewV4()), data.Name, data.Distance, &data.Date)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Insert data failed.",
				"data": err.Error(),
			})
			return
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
		err := c.BindJSON(&pagination)
		if err != nil {
			c.JSON(http.StatusOK, "Json deserialize failed.")
			return
		}
		rows, err := db.Query("SELECT INFO_ID, NAME, DISTANCE, DATE FROM info ORDER BY CREATE_TIME LIMIT ? , ?", pagination.Start, pagination.Size)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Query failed.",
				"data": err.Error(),
			})
			return
		}
		defer rows.Close()
		row := db.QueryRow("SELECT COUNT(*) FROM info;")
		total := 0
		row.Scan(&total)

		dataList := make([]Data, 0)
		for rows.Next() {
			var data = Data{}
			rows.Scan(&data.InfoId, &data.Name, &data.Distance, &data.Date)
			dataList = append(dataList, data)
		}
		c.JSON(http.StatusOK, PaginationResponse{Total: total, Data: dataList})
	}
}
