package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var userDB *sql.DB

// 定义用户的数据库结构
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 初始化数据库
func init() {
	_ = web.LoadAppConfig("ini", "conf/app2.conf")
	dbconn, _ := web.AppConfig.String("DBConn")
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	_ = db.Ping()
	userDB = db
}

func Close() {
	if userDB != nil {
		_ = userDB.Close()
	}

}

func AddUser(rec User) (User, error) {
	isql := "INSERT user SET ID=?,Name=?,Password=?"
	response := User{rec.ID, rec.Name, rec.Password}
	if userDB == nil {
		return response, errors.New("connect mysql failed")
	}
	stmt, err1 := userDB.Prepare(isql)
	if err1 != nil {
		fmt.Printf(err1.Error())
	}
	defer stmt.Close()
	_, err := stmt.Exec(rec.ID, rec.Name, rec.Password)
	if err == nil {
		response.Name = rec.Name
		response.Password = rec.Password
		return response, nil
	}

	return response, nil
}

func GetUser(name string) (User, error) {
	qsql := "SELECT * FROM user WHERE  name=?"
	var response User
	if name != "" {
		if userDB == nil {
			return response, errors.New("connect mysql failed")
		}
		stmt, _ := userDB.Prepare(qsql)
		rows, err := stmt.Query(name)
		defer rows.Close()
		if err != nil {
			return response, err
		}
		//遍历
		for rows.Next() {
			err = rows.Scan(&response.ID, &response.Name, &response.Password)
			if err != nil {
				return response, err
			}
		}
		return response, nil
	}
	return response, errors.New("Request failed!")
}
