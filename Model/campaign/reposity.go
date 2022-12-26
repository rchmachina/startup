package Models

import "gorm.io/gorm"

type Repository interface{
	//Create(user User) (User, error)
	FindAll() ([]Campaign,error )
	FindByUserID (userID int)([]Campaign,error)
	FindByID(id int)(Campaign,error)
	Save(campaign Campaign)	(Campaign, error)
	CheckName(name string)(Campaign, error)
}
type repository struct{
	db  *gorm.DB 

}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAll()([]Campaign,error){
	var campaigns []Campaign
	err := r.db.Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error
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

func (r *repository) FindByUserID(userID int)([]Campaign,error){
	var campaigns []Campaign
	err := r.db.Where("user_id = ?",userID).Preload("CampaignImage", "campaign_images.is_primary = 1").Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err!= nil{
		return campaigns, err
	}
	

	return campaigns ,nil

}


func (r *repository) FindByID(ID int)(Campaign,error){
	var campaigns Campaign
	err := r.db.Where("user_id = ?",ID).Preload("User").Preload("CampaignImage").Find(&campaigns).Error
	if err!= nil{
		return campaigns, err
	}
	

	return campaigns ,nil

}


func (r *repository) Save(campaign Campaign) (Campaign,error){
	err := r.db.Save(&campaign).Error
	if err !=nil{
		return campaign, err
	}
	return campaign, nil
}