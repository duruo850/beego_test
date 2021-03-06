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

func IsDbExist(dbName string) (bool, error) {
	qsql := "SELECT SCHEMA_NAME FROM SCHEMATA where SCHEMA_NAME = ?"
	var schemaName = ""
	stmt, err := schemaDB.Prepare(qsql)
	if err != nil {
		return false, err
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(dbName)
	if err != nil {
		return false, err
	}
	defer func() { _ = rows.Close() }()
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
	qsql := "create database if not exists " + dbName + "  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	stmt, err := schemaDB.Prepare(qsql)

	if err != nil {
		return false, err
	} else {
		defer func() { _ = stmt.Close() }()
		_, err := stmt.Exec()
		if err == nil {
			return true, nil
		}
		return false, err
	}
}
