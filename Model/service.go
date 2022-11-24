package Models

import (
	

	"golang.org/x/crypto/bcrypt"
)


type Service interface{
	RegisterUser(input RegisterUserInput) (User, error)

}

type service struct{
	repository Repository


}

func NewService(NewRepository Repository) *service{
	return &service{NewRepository}
}


func (s *service) RegisterUser(input RegisterUserInput)(User,error){

	user := User{}
	user.Name =  input.Name
	user.Email = input.Email
	user.Occupation= input.Occupation
	password,err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err!= nil{
		  return user, err
	}
	user.PasswordHash = string(password)
	if user.Role == ""{
		user.Role="user"
	}
	if user.AvatarFileName ==""{
		user.AvatarFileName="pictNewAvatar"
	}
	user.Token = input.Token
		
	newUser, err := s.repository.Save(user)
	if err!=nil{	
		return user, err
		}
	return newUser, nil
}
	


