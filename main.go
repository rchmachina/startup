package main

import (
	"campaign/Controller"
	// "campaign/Model"
	handler "campaign/handler"
	"fmt"

	
)

func main(){
	fmt.Print("test")
	// connection.Connect()
//	controller.Automigrates()
	// controller.MigrateUser()
//	

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	
	controller.Automigrates()
	//controller.Adduser()
	handler.RouterV1()
}


