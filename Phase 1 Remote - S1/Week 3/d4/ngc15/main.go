package main

import (
	"flag"
	"fmt"
	"log"
	"ngc15/config"
)

var name = flag.String("name", "", "Name Student")
var email = flag.String("email", "", "Email Student")

func main() {
	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Failed to connect db", err)
	}
	defer db.Close()

	flag.Parse()
	fmt.Println("Nama:", *name)
	fmt.Println("Email:", *email)
}
