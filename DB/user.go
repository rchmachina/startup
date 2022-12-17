package controller

import (
	Models "campaign/Model/user"
	"campaign/auth"
)

func Adduser(){
	db := Connect()
	userRepository := Models.NewRepository(db)
	user := Models.User{
		Name: "test simpan2",}
		userRepository.Create(user)

	userNewservice := Models.NewService(userRepository, auth.NewService())

	
	regUser := Models.RegisterUserInput{
		Name: "ini test lagi",
		Email:"test@gmail.com",
		Occupation: "programmernich",
		Password: "123554",

	}
	userNewservice.RegisterUser(regUser)


}

func GetUser(){
	
}
