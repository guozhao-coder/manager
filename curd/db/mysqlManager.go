package db

import (
	"smartroom_address/db/mysql"
	"smartroom_address/http_handler/bean"
)

func Mysql_Del(addr bean.DelAddr) error {
	return mysql.DB_del(addr)
}

func Mysql_Ins(addr bean.InsAddr) error{
	return mysql.DB_ins(addr)
}

func Mysql_GetAddr(addr bean.QueAddr) ([]*bean.InsAddr,error){
	addrs, e := mysql.DB_query(addr)
	return addrs,e
}

func Mysql_Upd(addr bean.InsAddr) error {
	return mysql.DB_update(addr)
}
