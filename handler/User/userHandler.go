package handler

import (
	controller "campaign/Controller"
	Models "campaign/Model"
	"campaign/auth"
	"campaign/helper"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//get user
func HandlerUser (c *gin.Context){
	db :=controller.Connect()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	var users []Models.User
	db.Find(&users)
	c.JSON(http.StatusOK,users)

	defer sqlDB.Close()

}




//post user
type userHandler struct{
	userService Models.Service
	authService auth.Service
}

func NewUserHandler(userService Models.Service, authService auth.Service) *userHandler{
	return &userHandler{userService, authService} 
}



func (h *userHandler) RegisterUser(c *gin.Context){
	var input Models.RegisterUserInput
	err := c.Bind(&input)
	if err !=nil{
		var errors []string

		for _, e := range err.(validator.ValidationErrors){
			errors = append(errors , e.Error())
			
		}

		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("register acc failed",http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return

	}
	


	inputdata,err:= h.userService.RegisterUser(input)
	if err!=nil{
		errorsMessage := gin.H{"errors": err.Error()}
		//
		c.JSON(http.StatusBadRequest,helper.APIResponse("fail",http.StatusBadRequest,"not updated",errorsMessage))
		return
	}
	token, err := h.authService.GenerateToken(int(inputdata.ID))
	if err != nil{
		return 
	}
	inputdata.Token = token
	


	   

	response:= helper.APIResponse("succes input data",http.StatusOK,"updated", inputdata)
	c.JSON(http.StatusOK, response)
}


func (h *userHandler) Login(c *gin.Context){

	var input Models.LoginInput

	err := c.Bind(&input)
	if err != nil{
		var errors []string

		for _, e := range err.(validator.ValidationErrors){
			errors = append(errors , e.Error())
			
		}

		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login failed",http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return

	}
	loggedInUser,err := h.userService.Login(input)
	
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("login failed",http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}
	token, err := h.authService.GenerateToken(int(loggedInUser.ID))
	if err != nil{
		return 
	}
	
	formatuser := helper.Formatuser(loggedInUser.Name, token)
	
	response:= helper.APIResponse("succes input data",http.StatusOK,"updated", formatuser)

	c.JSON(http.StatusOK, response)

}





func (h *userHandler) ChangeAvatar(c *gin.Context){

	var input Models.ChangeImageInput

	err := c.Bind(&input)
	if err != nil{
		var errors []string

		for _, e := range err.(validator.ValidationErrors){
			errors = append(errors , e.Error())
			
		}

		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login failed",http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return

	}
	file, err := c.FormFile("avatar")

	if err != nil{
		errorsMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest,helper.APIResponse("fail",http.StatusBadRequest,"please upload file",errorsMessage))
		return

	}
	var path string = fmt.Sprintf("images/" + file.Filename)
	pathformater := helper.Formaterword(path)
	c.SaveUploadedFile(file,pathformater)

	
	inputAvatar,err:= h.userService.ChangeAvatar(input.Token, path)
	if err!=nil{
		errorsMessage := gin.H{"errors": err.Error()}
		//
		c.JSON(http.StatusBadRequest,helper.APIResponse("fail",http.StatusBadRequest,"not updated",errorsMessage))
		return
	}

	IDstring := string(inputAvatar.ID) + "updated"
	response:= helper.APIResponse("Image updated",http.StatusOK,IDstring, inputAvatar)
	c.JSON(http.StatusOK, response)
}
