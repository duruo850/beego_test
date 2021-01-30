package mysql

import (
	"database/sql"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var utilDB *sql.DB = nil

func init() {
	initDB()
}

// 初始化数据库
func initDB() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")
	dbconn, _ := web.AppConfig.String("DBConn")

	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	err = db.Ping()
	if err != nil {
		return
	}
	utilDB = db
}

func ExecSql(querySql string) (bool, error) {
	if utilDB == nil {
		initDB()
	}
	stmt, err := utilDB.Prepare(querySql)
	if err != nil {
		return false, err
	} else {
		_, err := stmt.Exec()
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
}
