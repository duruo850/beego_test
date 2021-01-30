package mysql

import (
	"database/sql"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var schemaDB *sql.DB

// 初始化数据库
func init() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")
	dbconn, _ := web.AppConfig.String("DBConn")
	_db := strings.Split(strings.Split(dbconn, "/")[1], "?")[0]
	schemaDBConn := strings.Replace(dbconn, _db, "information_schema", -1)

	db, err := sql.Open("mysql", schemaDBConn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	_ = db.Ping()
	schemaDB = db
}

func IsDbExist(db_name string) (bool, error) {
	qsql := "SELECT SCHEMA_NAME FROM SCHEMATA where SCHEMA_NAME = ?"
	var schemaName string = ""
	stmt, err := schemaDB.Prepare(qsql)
	defer stmt.Close()
	if err != nil {
		return false, err
	}

	rows, err := stmt.Query(db_name)
	defer rows.Close()
	if err != nil {
		return false, err
	}
	//遍历
	for rows.Next() {
		err = rows.Scan(&schemaName)
		if err != nil {
			return false, err
		}
	}
	return schemaName != "", nil
}

func CreateDB(dbName string) (bool, error) {
	qsql := "create database if not exists " + dbName + " CHARACTER SET utf8 "
	stmt, err := schemaDB.Prepare(qsql)
	defer stmt.Close()
	if err != nil {
		return false, err
	} else {
		_, err := stmt.Exec()
		if err == nil {
			return true, nil
		}
		return false, err
	}
}
