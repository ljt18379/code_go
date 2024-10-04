package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func Write() {
	db := ConnMysql2()
	defer db.Close()
	_, err3 := db.Query("INSERT INTO baaschain VALUES('ssss')")
	if err3 != nil {
		log.Fatal(err3)
	}
}

func MysqldbQuery() {
	db := ConnMysql2()
	defer db.Close()
	rows, err := db.Query("SELECT module,pid,user,sum FROM poe")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var poe [4]interface{}
		err = rows.Scan(&poe[0], &poe[1], &poe[2], &poe[3])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(baas)
		fmt.Printf("poe: %s\n", poe)
	}
}

func ConnMysql() *sql.DB {
	// 数据源名
	driverName := "mysql"
	// 用户名root
	// 密码1234
	// tcp协议连接
	// 数据库地址
	// 数据库 wms
	dataSourceName := "root" + ":" + "123456" + "@" + "tcp" + "(" + "127.0.0.1:3306" + ")" + "/" + "mysql"
	db, err := sql.Open(driverName, dataSourceName)
	fmt.Printf("dataSourceName: %v\n", dataSourceName)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 连接测试
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	fmt.Println("sssss")
	return db
}

func ConnMysql2() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql")
	db.Ping()

	if err != nil {
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}

	//var version string
	//err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	return db
}
