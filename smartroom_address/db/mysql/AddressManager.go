package mysql

import (
	"database/sql"
	"errors"
	logs "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"smartroom_address/datasource"
	"smartroom_address/http_handler/bean"
)


/*
删除地点
 */
func DB_del(addr bean.DelAddr) error{
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()

	conn := datasource.DBConn
	begin, err := conn.Begin()
	if err != nil{
		logs.Error("DB_add error?:", err.Error())
		begin.Rollback()
		return err
	}

	//删除语句
	str := "delete from address where addr_id = ?"
	exec, err := begin.Exec(str,addr.AddrId)
	if err != nil{
		logs.Error("DB_delete exec err:",err)
		return err
	}
	if affected, _ := exec.RowsAffected(); affected == 1{
		begin.Commit()
		logs.Info("删除的行数：",affected)
		return nil
	}
	begin.Rollback()
	logs.Error("DB_delete affected rows is 0",err)
	return errors.New("DB_delete affected rows is 0")
}


/*
增加地点
 */
func DB_ins(addr bean.InsAddr) error{
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()
	conn := datasource.DBConn
	begin, err := conn.Begin()
	if err != nil{
		logs.Error("DB_add error?:", err.Error())
		begin.Rollback()
		return err
	}
	str := "insert into address" +
		" (addr_id , name , orderno , is_leaf , parent_id , src , addr_type , simple_name , nick_name , push_name , description) " +
		"values (?,?,?,?,?,?,?,?,?,?,?) "
	exec,err :=begin.Exec(str,addr.AddrId,addr.Name,addr.OrderNo,addr.IsLeaf,addr.ParentId,addr.Src,addr.AddrType,addr.SimpleName,addr.NickName,addr.PushName,addr.Description)
	if err != nil{
		logs.Error("DB_insert exec err:",err)
		return err
	}
	affected, err := exec.RowsAffected()
	if affected ==1{
		begin.Commit()
		logs.Info("增加的行数：",affected)
		return nil
	}else {
		logs.Error("DB_insert affected rows err:",err)
		begin.Rollback()
		return err
	}

}

//查询地点
func DB_query(addr bean.QueAddr) ([]*bean.InsAddr, error) {
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()
	conn := datasource.DBConn
	begin, err := conn.Begin()
	if err != nil{
		logs.Error("DB_query error?:", err.Error())
		begin.Rollback()
		return nil, err
	}
	str:= "select " +
		"a.addr_id," +
		"a.name," +
		"a.orderno," +
		"a.is_leaf," +
		"a.parent_id," +
		"a.addr_Type," +
		"a.simple_name," +
		"a.nick_name," +
		"a.push_name," +
		"a.description " +
		"from address a where name like CONCAT('%',?,'%',?,'%',?,'%',?,'%')"
	result,err:= begin.Query(str, addr.Center, addr.Model, addr.Floor, addr.Room)
	if err!= nil{
		logs.Error("query error",err)
		return nil,err
	}
	//定义一个封装地址的切片
	var addrs []*bean.InsAddr
	for result.Next(){
		var addrid sql.NullString
		var name sql.NullString
		var orderno sql.NullInt32
		var isleaf sql.NullInt32
		var parentid sql.NullString
		var addrtype sql.NullInt32
		var simplename sql.NullString
		var nickname sql.NullString
		var pushname sql.NullString
		var description sql.NullString
		err := result.Scan(&addrid, &name, &orderno, &isleaf, &parentid, &addrtype, &simplename, &nickname, &pushname, &description)
		if err != nil{
			logs.Error("result scan error:",err)
			return nil, err
		}
		
		a :=&bean.InsAddr{
			AddrId:      addrid.String,
			Name:        name.String,
			OrderNo:     orderno.Int32,
			IsLeaf:      isleaf.Int32,
			ParentId:    parentid.String,
			AddrType:    addrtype.Int32,
			SimpleName:  simplename.String,
			NickName:    nickname.String,
			PushName:    pushname.String,
			Description: description.String,
		}
		
		//追加到切片
		addrs = append(addrs,a)
	}
	return addrs,nil

}

/*
修改
 */
func DB_update(addr bean.InsAddr) error {
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()
	conn := datasource.DBConn
	begin, err := conn.Begin()
	if err != nil{
		logs.Error("DB_add error?:", err.Error())
		begin.Rollback()
		return err
	}
	str := "update address set name = ?, orderno = ?, is_leaf = ? ,parent_id = ?,addr_type = ?,simple_name = ?,nick_name = ?,push_name=?,description=? where addr_id=?"
	exec, err := begin.Exec(str,addr.Name, addr.OrderNo, addr.IsLeaf, addr.ParentId, addr.AddrType, addr.SimpleName, addr.NickName, addr.PushName, addr.Description, addr.AddrId)
	if err != nil{
		logs.Error("update exec err:",err)
		begin.Rollback()
		return err
	}
	affected, err := exec.RowsAffected()
	if affected == 0{
		begin.Rollback()
		logs.Error("DB_update affected rows is 0")
		return errors.New("DB_update affected rows is 0")
	}else {
		begin.Commit()
		logs.Info("修改的行数：",affected)
		return nil
	}

}