package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"golang-test/com/sql/model"
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

	stmt, err := db.Prepare("insert into order_info(name,age)values(?,?)")
	if err != nil {
		log.Println(err)
	}
	rs, err := stmt.Exec(name, age)
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

func findParam(pk int) int {
	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var name string
	var age int
	rows, err := db.Query("select name,age from order_info where id = ? ", pk)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &age)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name:", name, "age:", age)
	return 1
}

