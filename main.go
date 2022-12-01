package main

import (
	//"campaign/Controller"
	//controller "campaign/Controller"
	// models "campaign/Model"
	// "fmt"
	handler "campaign/handler"
	//"fmt"
)

func main(){

	// connection.Connect()
//	controller.Automigrates()
	// controller.MigrateUser()
//	

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	
	//controller.Automigrates()
	//controller.MigrateUser()
	//controller.Adduser()

	handler.RouterV1()
	// db := controller.Connect()
	// userRepository := models.NewRepository(db)
	// userService := models.NewService(userRepository)
	// //AddUserhandler := Handler.NewUserHandler(userService)

	// input := models.LoginInput{
	// 	Email : "rachlevi",
	// 	Password :"password",
	// }

	// User, err := userService.Login(input)
	// 	if err != nil{	
	// fmt.Print(err.Error())
	// }
	// fmt.Print(User.Email)
	// fmt.Print(User.Name)

}


