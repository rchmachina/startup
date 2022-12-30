package handler

import (
	DB "campaign/DB"
	//"campaign/helper"
	"campaign/middleware"
	//"strings"
	campaign_handler "campaign/handler/Campaign"
	models_campaign "campaign/Model/campaign"
	models_user "campaign/Model/user"
	"campaign/auth"
	UserHandler "campaign/handler/User"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (

	//user 
	db = DB.Connect()
	userRepository= models_user.NewRepository(db)
	userService = models_user.NewService(userRepository,auth.NewService())
	Userhandler = UserHandler.NewUserHandler(userService,auth.NewService())
	//campaign 

	campaignRepository= models_campaign.NewRepository(db)
	campaignService = models_campaign.NewService(campaignRepository)
	campaignhandler = campaign_handler.NewCampaignHandler(campaignService)


	
)


func RouterV1(){

	
	router :=gin.Default()
	router.Static("/images","./images")
	api :=router.Group("/api/v1")
	api.GET("/user",UserHandler.HandlerUser)
	api.GET("/", test)
	api.GET("/login", login)
	api.DELETE("/deletecampaign/:id",middleware.AuthMiddleware(auth.NewService(),userRepository),campaignhandler.DeleteCampaign)
 	api.GET("/campaign", campaignhandler.FindCampaignsAll)
	api.GET("/campaignsbyuser/:id", campaignhandler.FindCampaignsByUserId)
	api.PUT("/updatecampaigns/:id",  middleware.AuthMiddleware(auth.NewService(),userRepository),campaignhandler.Updatecampaign)
	api.POST("/CreateCampaign", middleware.AuthMiddleware(auth.NewService(),userRepository),campaignhandler.Createcampaign)
	api.POST("/Loginuser", Userhandler.Login)
	api.POST("/user",Userhandler.RegisterUser)
	api.DELETE("user")
	api.POST("/uploadAvatar", middleware.AuthMiddleware(auth.NewService(),userRepository),Userhandler.ChangeAvatar)
	router.Run()
}


func wellcome(c *gin.Context){
	array1 :=[2]int{123124,51251}
	c.JSON(http.StatusOK,array1)
}
func test(c *gin.Context){
	userByEmail, err := userRepository.FindByEmail("test@gmail.com")
	if err != nil{
		fmt.Print("error, not found")
	}

	
	c.JSON(http.StatusOK, userByEmail.Name)

}

func login(c *gin.Context){
	input := models_user.LoginInput{
		Email : "testing@gmail.com",
		Password :"12345",
	}

	User, err := userService.Login(input)
		if err == nil{	
			fmt.Print(User.Email)
			fmt.Print(User.Name)
		//	statusarr :=[]string{User.Email,User.Name}
			c.JSON(http.StatusOK, User.Email)
			}
		if err != nil{
			c.JSON(http.StatusBadRequest,"user or password is wrng")
		}
}