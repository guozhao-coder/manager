package config

import (
	"smartroom_address/util"
)

var dcm *DBConfig

//获取DBConfig的单例
func GetDBConfig() *DBConfig {
	if dcm == nil {
		dcm = &DBConfig{
			Mysql: new(SqlConfig),
		}
		util.ReadJsonFile("dev.json", dcm)

	}
	return dcm
}
