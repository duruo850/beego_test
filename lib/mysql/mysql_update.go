package mysql

import (
	"beego_test/lib"
	"bytes"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	_ "log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var versionDBConn string = ""
var versionDB *sql.DB
var curDBVersion int = 0
var minVersion int = 0
var maxVersion int = 0
var scriptMap map[int]string = make(map[int]string)
var _host string = ""
var _port string = ""
var _user string = ""
var _password string = ""
var _db string = ""

// DB_VERSION表格结构脚本
var dbVersionCreateTableSql = "CREATE TABLE `db_version` (`db_version` INT(11) NOT NULL DEFAULT '1' COMMENT '数据库版本') ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;"

// DB_VERSION表格初始数据脚本
var dbVersionInitDataSql = "INSERT  INTO `db_version`(`db_version`) VALUES (0);"

// 初始化数据库
func init() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")

	versionDBConn, _ = web.AppConfig.String("DBConn")
	dbSettings := strings.Split(versionDBConn, "@")
	dbUserSettings := strings.Split(dbSettings[0], ":")
	dbHostSettings := strings.Split(strings.Split(strings.Split(dbSettings[1], "(")[1], ")")[0], ":")
	_host = dbHostSettings[0]
	_port = dbHostSettings[1]
	_user = dbUserSettings[0]
	_password = dbUserSettings[1]
	_db = strings.Split(strings.Split(versionDBConn, "/")[1], "?")[0]

	ensureVersionDB()

	// 加载脚本
	loadScript()
}

func ensureVersionDB() {
	if versionDB == nil {
		versionDB = initDB(versionDBConn)
	}
}

func Update() {
	isExist, err := IsDbExist(_db)
	if err != nil {
		fmt.Printf("IsDbExist: %s failed!!!", _db)
	}
	if isExist == false {
		_, err = CreateDB(_db)
		if err != nil {
			fmt.Printf("CreateDB: %s failed!!!, err: %s\n", _db, err)
			os.Exit(1)
		} else {
			fmt.Printf("CreateDB: %s success!!!\n", _db)
		}
		_, err = ExecSql(dbVersionCreateTableSql)
		if err != nil {
			fmt.Printf("ExecSql: %s failed!!!, err:%s\n", dbVersionCreateTableSql, err)
			os.Exit(1)
		} else {
			fmt.Printf("ExecSql: %s success!!!\n", dbVersionCreateTableSql)
		}
		_, err = ExecSql(dbVersionInitDataSql)
		if err != nil {
			fmt.Printf("ExecSql: %s failed!!!, err:%s\n", dbVersionInitDataSql, err)
			os.Exit(1)
		} else {
			fmt.Printf("ExecSql: %s success!!!\n", dbVersionInitDataSql)
		}
	}
}

func loadScript() {
	curWD, _ := os.Getwd()

	// 加载脚本
	versionDir := curWD + "/db/mysql_update/version"
	mysqlScripts, _ := ioutil.ReadDir(versionDir)
	var versionLS []int

	for _, mysqlFile := range mysqlScripts {
		arr := strings.Split(mysqlFile.Name(), ".")
		version, _ := strconv.Atoi(arr[0])
		ftype := arr[1]
		if ftype != "sql" {
			continue
		}
		versionLS = append(versionLS, version)
		scriptMap[version] = versionDir + "/" + mysqlFile.Name()
	}

	minVersion, maxVersion = lib.FindMinMax(versionLS)
}

// 执行sql脚本
func execScript(script string) {
	var b bytes.Buffer
	b.WriteString("mysql -h ")
	b.WriteString(_host)
	b.WriteString(" -P")
	b.WriteString(_port)
	b.WriteString(" -u")
	b.WriteString(_user)
	b.WriteString(" -p")
	b.WriteString(_password)
	b.WriteString(" ")
	b.WriteString(_db)
	b.WriteString(" < ")
	b.WriteString(script)
	cmd := exec.Command("/bin/bash", "-c", b.String()) //不加第一个第二个参数会报错

	stdout, _ := cmd.StdoutPipe() //创建输出管道
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Printf("error: %v\n", err)
	} else {
		log.Printf(" <db:%s> successful exec sql file = %s", _db, script)
	}
}

// 执行sql脚本
//func get_cur_db_version(script string) {
//	qsql := "select db_version from db_version"
//	if userDB == nil {
//			return response, errors.New("connect mysql failed")
//		}
//	stmt, _ := userDB.Prepare(qsql)
//	rows, err := stmt.Query(name)
//	defer rows.Close()
//	if err != nil {
//		return response, err
//	}
//	//遍历
//	for rows.Next() {
//		rows
//		err = rows.Scan(&response.ID, &response.Name, &response.Password)
//		if err != nil {
//			return response, err
//		}
//	}
//	return response, nil
//}

//@property
//def cur_db_version(self):
//"""
//获得数据库版本
//"""
//if self.__db_version is not None:
//return self.__db_version
//
//sql_cmd = 'select db_version from db_version'
//ret = self.__mysql_clt.query(sql_cmd)
//if not ret.success:
//# 初始版本是0
//return 0
//self.__db_version = ret.first()['db_version']
//return self.__db_version
