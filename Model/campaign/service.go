package Models

import (
	
	"errors"
)


type Service interface{
	FindCampaignsByuserId(UserID int) ([]Campaign, error)
	FindCampaignByid(ID int) (Campaign, error)
}


type service struct{
	repository Repository
}

func NewService(NewRepository Repository) *service {
	return &service{NewRepository}
}


func (s * service) FindCampaignsByuserId(UserID int)([]Campaign, error){
	var campaign []Campaign
	if UserID ==0{
		
	campaigns, err := s.repository.FindAll()
	if err != nil{
		return campaigns, err
		}
		
	return campaigns,nil	
	}
	campaign, err := s.repository.FindByUserID(UserID)
	if err !=nil{
		return campaign,errors.New("error")
	}

	return campaign,nil	
	}


func (s * service) FindCampaignByid(ID int)(Campaign, error){
	var campaign Campaign
		
	campaign, err := s.repository.FindByID(ID)
	if err != nil{
		return campaign, err
		}
		
	return campaign,nil	
	}
	
	
	

	
