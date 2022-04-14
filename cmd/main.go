package main

import (
	"benching/db"
	"fmt"
	"log"
)

func main() {
	dbLocal, err := db.NewDB("members.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbLocal.Close()

	err = dbLocal.AddNewMember()
	if err != nil {
		log.Fatal(err)
	}

	members, err := dbLocal.GetAllMember()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(members)
}
