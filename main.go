package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username  string
	password  string
	mysqlname string
	mysqlpass string
)

const ()

func init() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\033[H\033[2J")
		fmt.Printf(`

==============================
	bye %s
==============================

	`, username)
		os.Exit(0)
	}()
}

func main() {
	UserScreen()
}

func UseMysl() {
	fmt.Println("please enter your mysql username")
	fmt.Scanln(&mysqlname)
	fmt.Println("please enter your mysql password")
	fmt.Scanln(&mysqlpass)
	_, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", mysqlname, mysqlpass))
	if err != nil {
		fmt.Println("have some problems !")
	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println(
			`
==============================
talk me what do you want to do

input
cDB : create database
t 	: show 	 tables

==============================
			`)
	}
}

// func Mysql() {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", name, pass))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	var dbname string = "hellomysql"
// 	// _, err = db.Exec("CREATE DATABASE " + dbname)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	_, err = db.Exec("USE " + dbname)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//TODO:未來有機會想出怎麼把TABLE讓使用者自定義寫進去
// 	res, _ := db.Query("SHOW TABLES")
// 	var tableList string
// 	for res.Next() {
// 		res.Scan(&tableList)
// 		fmt.Println(tableList)
// 	}
// }

func UserScreen() {
	var loginFlag bool = false
	fmt.Print(
		`
==============================
HI USER WELCOME TO PETER SQL 
==============================

INPUT YOUR USERNAME
`)
	fmt.Scanf("%s\n", &username)
	fmt.Print("\nINPUT YOUR USERNAME\n")
	fmt.Scanf("%s", &password)
	file, err := os.OpenFile("user.txt", os.O_RDWR, 777)
	_, err = file.Read([]byte("peter,password"))
	if err == io.EOF {
		fmt.Println("have some problems")
		return
	} else {
		loginFlag = true
	}

	if loginFlag {
		fmt.Printf("\nWait to login ... \n")
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Printf(`
==============================
HI %s NICE TO MEET YOU
==============================
`, username)
		fmt.Println()
		UseMysl()
	}
}
