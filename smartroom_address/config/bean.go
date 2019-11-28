package config

//DB的配置参数
type DBConfig struct {
	DBType string        `json:"dbType"` //数据库类型
	Mysql  *SqlConfig    `json:"mysql"`  //mysql数据库

}


//sql数据库配置参数
type SqlConfig struct {
	Driver  string `json:"driver"`  //数据库驱动
	Host    string `json:"host"`    //数据库地址
	Port    int    `json:"port"`    //端口号
	Account string `json:"account"` //账号
	Pwd     string `json:"pwd"`     //密码
	DBName  string `json:"dbname"`  //数据库名
}



