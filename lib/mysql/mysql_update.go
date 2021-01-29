package mysql

import (
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

// 初始化数据库
func init() {
	_ = web.LoadAppConfig("ini", "conf/app2.conf")

	dbconn, _ := web.AppConfig.String("DBConn")
	// DBConn="root:123456@tcp(localhost:3306)/user?charset=utf8"
	dbSettings := strings.Split(dbconn, "@")
	dbUserSettings := strings.Split(dbSettings[0], ":")
	dbHostSettings := strings.Split(strings.Split(strings.Split(dbSettings[1], "(")[1], ")")[0], ":")
	_host = dbHostSettings[0]
	_port = dbHostSettings[1]
	_user = dbUserSettings[0]
	_password = dbUserSettings[1]
	_db = strings.Split(strings.Split(dbconn, "/")[1], "?")[0]

	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	_ = db.Ping()
	versionDB = db

	loadScript()
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
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

		fmt.Println(mysqlFile.Name())
	}

	minVersion, maxVersion = findMinAndMax(versionLS)
	fmt.Printf("minVersion %d, maxVersion: %d\n", minVersion, maxVersion)
}

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
