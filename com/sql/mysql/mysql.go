package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"golang-test/com/sql/model"
	"time"
)

func findByPk(pk int) int {
	var num int = 0
	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("select * from order_info where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(pk).Scan(&num)
	if err != nil {
		panic(err.Error())
	}
	return num
}


//func getDb() sql.DB{
//	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
//	if err != nil {
//		panic(err.Error())
//	}
//	return db;
//}

func insert(name string, age int) interface{}  {
//func insert(name string, age int) model.OrderInfoResponse  {
	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into order_info(name,age,create_time,update_time,version)values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}

	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	//rs, err := stmt.Exec(name, age, time.Now(), time.Now(), 0)
	rs, err := stmt.Exec(name, age, datetime, datetime, 0)
	if err != nil {
		log.Println(err)
	}
	//var id int
	//var affect int
	//我们可以获得插入的id
	id, err := rs.LastInsertId()
	//可以获得影响行数
	affect, err := rs.RowsAffected()
	fmt.Println("id:", id, "affect:", affect)
	//var orderInfo *model.OrderInfoResponse = new(model.OrderInfoResponse)
	//orderInfo.id = id
	//orderInfo.affect = affect
	 orderInfoResponse  := &model.OrderInfoResponse{Id:id,Affect:affect}
	return orderInfoResponse
}

func findParam(pk int) interface{} {
	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var id int
	var name string
	var age int
	var create_time time.Time
	var update_time time.Time
	var version int
	rows, err := db.Query("select * from order_info where id = ? ", pk)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &create_time, &update_time, &version)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	orderInfo := &model.OrderInfo{Id:id,Name:name,Age:age,CreateTime:create_time,UpdateTime:update_time,Version:version}

	fmt.Println("orderInfo:", orderInfo)
	return orderInfo
}

