package Models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)


type Service interface{
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User,error)
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
	
func(s *service) Login(input LoginInput) (User,error){
	email := input.Email
	password := input.Password

	user, err :=s.repository.FindByEmail(email)
	if err != nil{
		return user, errors.New("email not found")
	}
	if user.ID==0{
		return user, errors.New("not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err !=nil{
		return user, errors.New("password wrong")
		}
	return user,nil
}

