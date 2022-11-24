package handler


import (
	Handler "campaign/handler/folderHandler"
	"net/http"
	"campaign/Controller"
	models "campaign/Model"
	"github.com/gin-gonic/gin"
	
)

var (


	db = controller.Connect()
	userRepository= models.NewRepository(db)
	userService = models.NewService(userRepository)
	AddUserhandler = Handler.NewUserHandler(userService)
	

)






func RouterV1(){

	
	router :=gin.Default()
	api :=router.Group("/api/v1")
	api.GET("/user",Handler.HandlerUser)
	api.GET("/", wellcome)
	router.POST("/user",AddUserhandler.RegisterUser)
	router.Run()
}


func wellcome(c *gin.Context){
	array1 :=[2]int{123124,51251}
	c.JSON(http.StatusOK,array1)
}
