package mysql

import (
	"database/sql"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var utilDBConn string = ""
var utilDB *sql.DB = nil

func init() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")
	utilDBConn, _ = web.AppConfig.String("DBConn")
	utilDB = initDB(utilDBConn)
}

func ensureUtilDB() {
	if utilDB == nil {
		utilDB = initDB(utilDBConn)
	}
}

// 初始化数据库
func initDB(dbconn string) *sql.DB {
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return nil
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	err = db.Ping()
	if err != nil {
		return nil
	}
	return db
}

func ExecSql(querySql string) (bool, error) {
	ensureUtilDB()
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
