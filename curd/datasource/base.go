package datasource

import (
	"database/sql"
	"fmt"
	logs "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"smartroom_address/config"
)

var DBType int = 0
var DBConn *sql.DB

const (
	_ = iota
	MYSQL
	TIDB
)
const (
	TIDBSTR  = "TIDB"
	MYSQLSTR = "MYSQL"
)

func init() {
	dbc := config.GetDBConfig()
	var driverName string = ""
	var url string = ""

	//获取MySQL客户端
	if dbc.DBType == TIDBSTR {
		DBType = TIDB
		driverName = dbc.Mysql.Driver
		url = fmt.Sprintf("%s:%s@/%s", dbc.Mysql.Account, dbc.Mysql.Pwd, dbc.Mysql.DBName)

	}
	logs.Info("drivername:", driverName)
	dbconn, err := sql.Open(driverName, url)
	if err != nil {

		logs.Error("数据库连接失败", err.Error())
		os.Exit(12)
	}
	err = dbconn.Ping()
	if err != nil {
		fmt.Println("dbconn is nil", err.Error())
		os.Exit(2)
	} else {
		logs.Info("db conn success")
		DBConn = dbconn
		DBConn.SetMaxIdleConns(50)
		DBConn.SetMaxOpenConns(512)
	}
}
