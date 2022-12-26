package main

import (
	database "campaign/DB"


	"fmt"
	//
	auth "campaign/auth"

	handler "campaign/handler"
)

var (

	db = database.Connect()

	//campaignService = model_campaigns.NewService(campaignRepository)
	//campaignhandler = campaign_handler.NewCampaignHandler(campaignService)


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
	
	//database.Automigrates()
	//controller.MigrateUser()
	//controller.Adduser()
	//userService.SaveAvatar(4,"image/1")
	authservice := auth.NewService()
	
	fmt.Println(authservice.GenerateToken(33))

	//database.MigrateCampaign()

	// campaigns, err := campaignRepository.FindByUserID(int(6969))
	// 	if err!=nil{
	// 		fmt.Println("there something wrong")
	// 	}

	// for _, campaign :=range campaigns{
	// 	fmt.Println(campaign.Name)
	// 	fmt.Println(campaign.UserID)
	// 	if len(campaign.CampaignImage)>0{
	// 		fmt.Println(campaign.CampaignImage[0].FileName)
	// 	}
	// }

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


