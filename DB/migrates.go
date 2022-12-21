package controller

import (
	//model_user "campaign/Model/user"
	models_campaign "campaign/Model/campaign"
	model_user "campaign/Model/user"

	"log"
)

func Automigrates(){


	db := Connect()


	DB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	//db.AutoMigrate(model_user.User{},)
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.User{})
	// db.AutoMigrate(model_campaign.CampaignImage{})
	// db.AutoMigrate(model_campaign.Campaign{},)
	defer DB.Close()


	
}

func MigrateCampaign(){
// 	// var insertCampaign = model_campaign.Campaign{

// 	// 	ID:69, UserID: 69, Name:"hahahaha", ShortDescription:"fafifefo", Description : "faifhainauifaiuwnbIOANF",
// 	// 	Perks:"this is perks", BackerCount: 1000, GoalAmount:100000000,CurrentAmount:1000000,Slug:"this is slug",
// 	// 	ImageUrl:"this is image url",CampaignImage: []model_campaign.CampaignImage{
// 	// 		{CampaignID:100,FileName:"fafufefo.jpeg",IsPrimary:1},
// 	// 		{CampaignID:100,FileName:"fafufefo1.jpeg",IsPrimary:0}, 
// 	// 		{CampaignID:100,FileName:"fafufefo1.jpeg",IsPrimary:0},
// 	// 	},
// 	// }
	db:= Connect()
	sqlDB, err := db.DB()
		if err != nil {
		log.Fatalln(err)
	 }
	// variation := models_campaign.Campaign{

	// 	UserID: 666, Name:"hahahaha", ShortDescription:"fafifefo", Description : "faifhainauifaiuwnbIOANF",
	// 	Perks:"this , is , perks", BackerCount: 1000, GoalAmount:100000000,CurrentAmount:1000000,Slug:"this is slug",
	// 	ImageUrl:"this is image url",CampaignImage: []models_campaign.CampaignImage{
			
	// 		{CampaignID:666,FileName:"fafuf331efo1.jpeg",IsPrimary:0}, 
	// 		{CampaignID:666,FileName:"fafufefo3442.jpeg",IsPrimary:1},
	// 		{CampaignID:66,FileName:"faf.jpeg",IsPrimary:0},
	// 	},


	// }
	campaign := models_campaign.Campaign{}
	db.Model(&campaign).Where("id = ?", 666).Update("perks", "hello , this, is, slug")
	
	defer sqlDB.Close()
	//db.Create(&variation)

}


func MigrateUser(){
	// var p3 = []Models.User{
	// 	{Name: "fafofe", Occupation: "test", Email: "test1@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"},
	// 	{Name: "EZZZZ", Occupation: "12313", Email: "test2@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"},
		
	// }

	variation :=model_user.User{ 
		ID:69 ,Name: "faf", Occupation: "test", Email: "test1@gmail.com", PasswordHash: "1111", AvatarFileName: "test.jpeg"}
		
	db := Connect()


	DB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	
	defer DB.Close()
	// db.Create(&p3)
	db.Create(&variation)
}