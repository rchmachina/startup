package Models

import (

	"gorm.io/gorm"
)


type Repository interface{
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(ID int)(User, error)
	Update(user User) (User,error)
	FindByToken(token string) (User, error)
	
	
}
type repository struct{
	db  *gorm.DB

}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}


func (r *repository) Create(user User) (User,error){
	err := r.db.Create(&user).Error
	if err !=nil{
		return user, err
	}
	

	return user, nil
}

func (r *repository) Update(user User) (User,error){
	err := r.db.Save(&user).Error
	if err !=nil{
		return user, err
	}
	return user, nil
}
func (r *repository) CreateToken(id int)(User, error){
	var user User
	err := r.db.Where("ID=? ", id).Find(&user).Error
	if err !=nil{
		return user,err
	}
	return user, nil

} 



func (r *repository) FindByEmail(email string)(User, error){
	var user User
	err := r.db.Where("email =? ", email).Find(&user).Error
	if err !=nil{
		return user,err
	}
	return user, nil

} 

func (r *repository) FindById(ID int)(User, error){
	var user User
	err := r.db.Where("ID=? ", ID).Find(&user).Error
	if err !=nil{
		return user,err
	}
	return user, nil

} 

func (r *repository) FindByToken(token string)(User, error){
	var user User
	err := r.db.Where("Token=? ", token).Find(&user).Error
	if err !=nil{
		return user,err
	}
	return user, nil

} 


