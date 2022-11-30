package handler

import (
	"campaign/Controller"
	models "campaign/Model"
	Handler "campaign/handler/folderHandler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (


	db = controller.Connect()
	userRepository= models.NewRepository(db)
	userService = models.NewService(userRepository)
	Userhandler = Handler.NewUserHandler(userService)


)






func RouterV1(){

	
	router :=gin.Default()
	api :=router.Group("/api/v1")
	api.GET("/user",Handler.HandlerUser)
	api.GET("/", test)
	api.GET("/login", login)
	api.POST("/login", Userhandler.Login)
	router.POST("/user",Userhandler.RegisterUser)
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
	input := models.LoginInput{
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