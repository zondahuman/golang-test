package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"golang-test/com/persist/model"
	"time"
	"errors"
)

func findByPk(pk int) int {
	var num int = 0
	db := getDb()
	defer db.Close()
	//var id int
	//id, err = db.Get(&id, "SELECT count(*) FROM place")
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

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return db;
}

func insert(name string, age int) interface{} {
	db := getDb()
	defer db.Close()

	stmt, err := db.Prepare("insert into order_info(name,age,create_time,update_time,version)values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}

	var datetime = time.Now()
	//datetime.Format(time.RFC3339)
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
	orderInfoResponse := &model.OrderInfoResponse{Id:id, Affect:affect}
	return orderInfoResponse
}

func findParam(pk int) interface{} {
	db := getDb()
	defer db.Close()

	var id int
	var name string
	var age int
	var create_time string
	var update_time string
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
	DefaultTimeLoc := time.Local
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", create_time, DefaultTimeLoc)
	updateTime, _ := time.Parse("2006-01-02 15:04:05", update_time)
	orderInfo := &model.OrderInfo{Id:id, Name:name, Age:age, CreateTime:createTime, UpdateTime:updateTime, Version:version}

	fmt.Println("orderInfo:", orderInfo)
	return orderInfo
}

func updateAge(age int, id int) {
	if r := recover(); r != nil {

		db := getDb()
		defer db.Close()
		//方式4 update
		start := time.Now()
		tx, _ := db.Begin()
		tx.Exec("update order_info set age=? where id=?", age, id)
		tx.Commit()
		end := time.Now()
		fmt.Println(" update total time:", end.Sub(start).Seconds())

		   switch x := r.(type) {
                     case string:
                            err = errors.New(x)
                     case error:
                            err = x
                     default:
                            err = errors.New("Unknow panic")
                     }
	}

}

