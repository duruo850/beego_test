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

var versionDBConn = ""
var versionDB *sql.DB
var curDBVersion = 0
var maxVersion = 0
var scriptMap = make(map[int]string)
var _host = ""
var _port = ""
var _user = ""
var _password = ""
var _db = ""

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

	_, maxVersion = lib.FindMinMax(versionLS)
}

// 执行sql脚本
func execScript(script string) error {
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

	//不加第一个第二个参数会报错
	cmd := exec.Command("/bin/bash", "-c", b.String())
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if err := cmd.Run(); err != nil {
		log.Printf("<db:%s> failed exec sql file:%s error: %s\n", _db, script, w)
		return err
	} else {
		log.Printf(" <db:%s> successful exec sql file = %s", _db, script)
	}
	return nil
}

// 查询当前版本
func getCurVersion() (int, error) {
	ensureVersionDB()
	qsql := "select db_version from db_version"
	stmt, err := versionDB.Prepare(qsql)
	if err != nil {
		return 0, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()
	curVersion := 0
	//遍历
	for rows.Next() {
		err = rows.Scan(&curVersion)
		if err != nil {
			return 0, err
		}
	}
	return curVersion, nil
}

// 设置当前版本
func setCurVersion(version int) error {
	ensureVersionDB()
	qsql := "update db_version set db_version = ?"
	stmt, err := versionDB.Prepare(qsql)
	if err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	_, err = stmt.Exec(version)
	if err != nil {
		return err
	}
	return nil
}

// 设置当前版本
func updateToVersion(version int) error {
	ensureVersionDB()
	toExecScript := scriptMap[version]
	err := execScript(toExecScript)
	if err != nil {
		fmt.Printf("execScript:%s failed!!!, err:%s\n", toExecScript, err)
		os.Exit(1)
		return err
	}

	err = setCurVersion(version)
	if err != nil {
		fmt.Printf("setCurVersion:%d failed!!!, err:%s\n", version, err)
		os.Exit(1)
	}
	return nil
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

	curDBVersion, err = getCurVersion()
	if err != nil {
		fmt.Printf("getCurVersion: failed!!!, err:%s\n", err)
		os.Exit(1)
	}
	if curDBVersion > maxVersion {
		fmt.Printf("curDBVersion:%d bigger than: maxVersion:%d\n", curDBVersion, maxVersion)
		os.Exit(1)
	}

	updateVersion := curDBVersion + 1
	for {
		if updateVersion > maxVersion {
			break
		}
		err := updateToVersion(updateVersion)
		if err != nil {
			fmt.Printf("updateToVersion:%d failed!!!, err:%s\n", updateVersion, err)
			os.Exit(1)
		}
		updateVersion++
	}
}
