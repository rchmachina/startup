package Models

import (
	"errors"

	"gorm.io/gorm"
)
		


type Repository interface{
	//Create(user User) (User, error)
	FindAll() ([]Campaign,error )
	FindByUserID (userID int)([]Campaign,error)
	Update(campaign Campaign)(Campaign, error)
	Save(campaign Campaign)	(Campaign, error)
	CheckName(name string)(Campaign, error)
	checkImage(campaignID int)([]CampaignImage,error)
	FindByID (ID int)(Campaign,error)
	CreateImage (campaign_image CampaignImage)(CampaignImage,error)
	DeleteCampaign (ID int)(error)
}
type repository struct{
	db  *gorm.DB 

}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}
func (r*repository) DeleteCampaign (ID int)(error){

	err := r.db.Delete(&Campaign{}, ID).Error
	if err !=nil{
		return errors.New("there is no data in here ")
	}
	
	return nil

}



func (r*repository) CreateImage (campaign_image CampaignImage)(CampaignImage,error){
	err := r.db.Create(&campaign_image).Error
	if err !=nil{
		return campaign_image, err
	}
	return campaign_image, nil
}



func (r *repository) FindAll()([]Campaign,error){
	var campaigns []Campaign
	err := r.db.Preload("CampaignImage", "campaign_images.is_primary = 1").Preload("User").Find(&campaigns).Error
	if err!= nil{
		return campaigns, err
	}

	return campaigns,nil

}

func (r *repository) CheckName(name string)(Campaign, error){
	var campaign Campaign
	err := r.db.Where("Name=? ", name).Find(&campaign).Error
	if err !=nil{
		return campaign,err
	}
	return campaign, nil

} 



func (r *repository) checkImage(campaignID int)([]CampaignImage,error){
	var CampaignImage []CampaignImage
	err := r.db.Where("campaign_id = ?",campaignID).Find(&CampaignImage).Error
	if err !=nil{
		return CampaignImage,err
	}
	return CampaignImage, nil

} 


func (r *repository) FindByUserID(userID int)([]Campaign,error){
	var campaigns []Campaign
	err := r.db.Where("user_id = ?",userID).Preload("User").Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err!= nil{
		return campaigns, err
	}
	

	return campaigns ,nil

}
func (r *repository) FindByID(ID int)(Campaign,error){
	var campaigns Campaign
	err := r.db.Preload("User").Preload("CampaignImage").Where("id = ?",ID).Find(&campaigns).Error
	if err!= nil{
		return campaigns, err
	}
	

	return campaigns ,nil

}



func (r *repository) Save(campaign Campaign) (Campaign,error){
	err := r.db.Create(&campaign).Error
	if err !=nil{
		return campaign, err
	}
	return campaign, nil
}



func (r *repository) Update(campaign Campaign) (Campaign,error){
	err := r.db.Save(&campaign).Error
	if err !=nil{
		return campaign, err
	}
	return campaign, nil
}