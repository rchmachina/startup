package controller

import (
	Models "campaign/Model"
	"log"
)

func Automigrates(){


	db := Connect()


	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(Models.User{},)
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.User{})

	defer sqlDB.Close()


	
}


func MigrateUser(){
	var p3 = []Models.User{
		{Name: "fafofe", Occupation: "test", Email: "test1@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"},
		{Name: "EZZZZ", Occupation: "12313", Email: "test2@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"},
		
	}

	variation :=Models.User{
		Name: "faf", Occupation: "test", Email: "test1@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"}
		
	db := Connect()


	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	
	defer sqlDB.Close()
	db.Create(&p3)
	db.Create(&variation)
}