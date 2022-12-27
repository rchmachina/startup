package Models

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)


type Service interface{
	FindCampaignsByuserId(UserID int) ([]Campaign, error)
	CreateCampaign(input CreateCampaign) (Campaign, error)

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


	
	


	
func (s *service) CreateCampaign(input CreateCampaign) (Campaign, error) {
	campaign := Campaign{}
	
	campaign.Name = input.Name
	checkname, err:= s.repository.CheckName(campaign.Name)
	if err != nil{
		return checkname,err
	}
	if checkname.ID != 0{
		return checkname, errors.New("nama campaign already exist")

	}
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = int(input.User.ID)
	campaign.Perks = input.Perks
	varslug := fmt.Sprintf("%s+%d", campaign.Name, campaign.UserID)
	campaign.Slug = slug.Make(varslug)
	NewCampaign, err := s.repository.Save(campaign)
	if err != nil{
		return NewCampaign,err
	}
	return NewCampaign,nil

} 