package service

import (
	"smartroom_address/db"
	"smartroom_address/http_handler/bean"
)
//删除地点
func Delete_Address(addr bean.DelAddr) error{
	return db.Mysql_Del(addr)
}
//新增地点
func Insert_Address(addr bean.InsAddr) error{
	return db.Mysql_Ins(addr)
}

//模糊查询
func Query_Address(addr bean.QueAddr) ([]*bean.InsAddr,error){
	return db.Mysql_GetAddr(addr)
}

//编辑
func Update_Address(addr bean.InsAddr) error{
	return db.Mysql_Upd(addr)
}


