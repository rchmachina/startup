package Models

import (
	
	"errors"
	//"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	//SaveAvatar(ID int, fileLocation string) (User, error)
	//CheckEmail(user User) (bool, error)
	ChangeAvatar(token string, file string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(NewRepository Repository) *service {
	return &service{NewRepository}
}

// buat func samain user struct dengan image structnya
//ambil repositi dari find by email
//jika dari emailnya ada ambil IDnya
//save ke repositi yg ada dbnya

func (s *service) ChangeAvatar(token string, path string) (User, error) {

	finduser, err := s.repository.FindByToken(token)
	if err != nil {
		return finduser, errors.New("token not found")
	}

	user := User{ID: finduser.ID, Role: finduser.Role, Name: finduser.Role, PasswordHash: finduser.PasswordHash,
		AvatarFileName: path, Occupation: finduser.Occupation, Token: finduser.Token, CreatedAt: finduser.CreatedAt, Email: finduser.Email}

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}
	return updateUser, nil

}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email

	user.Occupation = input.Occupation

	checkemail, _ := s.repository.FindByEmail(user.Email)
	// if err != nil{
	// 	return checkemail, errors.New("email already exist")
	// }
	if checkemail.ID != 0 {
		return checkemail, errors.New("email already exist")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(password)
	if user.Role == "" {
		user.Role = "user"
	}
	if user.AvatarFileName == "" {
		user.AvatarFileName = "pictNewAvatar"
	}




	UpdatedUser, err := s.repository.Create(user)
	if err != nil {
		return user, err
	}




	return UpdatedUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, errors.New("email not found")
	}
	if user.ID == 0 {
		return user, errors.New("not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("password wrong")
	}

	return user, nil
}

// func (s *service) SaveAvatar(ID int, fileLocation string)(User, error){

// 	user, err := s.repository.FindById(ID)
// 	if err != nil{
// 		return user, err

// 	}
// 	user.AvatarFileName = fileLocation

// 	s.repository.Update(user)
// 	updatedUser , err := s.repository.Update(user)
// 	if err != nil{
// 		return updatedUser, err
// 	}
// 	return updatedUser,nil
// }



