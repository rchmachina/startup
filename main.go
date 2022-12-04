package main

import (

	// controller "campaign/Controller"
	// models "campaign/Model"
	// "fmt"
	handler "campaign/handler"
	//UserHandler "campaign/handler/User"
	//"fmt"
)

// var (


// 	db = controller.Connect()
// 	userRepository= models.NewRepository(db)
// 	userService = models.NewService(userRepository)
// 	Userhandler = UserHandler.NewUserHandler(userService)

	
// )


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
	//userService.SaveAvatar(4,"image/1")
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


