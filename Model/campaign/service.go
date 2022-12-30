package Models

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)


type Service interface{
	FindCampaignsByUserId(UserID int) ([]Campaign, error)
	CreateCampaign(input CreateCampaign) (Campaign, error)
	Updatecampaign (inputID int, input CreateCampaign) (Campaign, error)
	FindById(ID int)(Campaign, error)
	UploadCampaignImagePrimary(input CampaignImage) (Campaign, error)
	DeleteCampaign(input int )(error)

}


type service struct{
	repository Repository

}

func NewService(NewRepository Repository) *service {
	return &service{NewRepository}
}
func (s * service)DeleteCampaign(input int )(error){
	err := s.repository.DeleteCampaign(input)
	if err != nil{
		return err
	}
	return nil
}


func (s * service) FindCampaignsByUserId(UserID int)([]Campaign, error){
	if UserID == 0{


		campaign, err := s.repository.FindAll()
		if err != nil{
			return campaign, err
		}
		return  campaign,nil

	}
	campaign, err := s.repository.FindByUserID(UserID)
	if err !=nil{
		return campaign,errors.New("error")
	}
	return campaign,nil	

	}	

	

func (s * service) FindById(ID int)(Campaign, error){
	var campaign Campaign

	campaign, err := s.repository.FindByID(ID)
	if err !=nil{
		return campaign,errors.New("error")
	}

	return campaign,nil	
	}


func (s * service) Updatecampaign(inputID int, input CreateCampaign)(Campaign, error){
	var campaign Campaign
	campaign, err := s.repository.FindByID(inputID)
	if err !=nil{
		return campaign ,errors.New("data tidak ditemukan")
	}
	campaign.ID = inputID
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Name = input.Name

	campaign.Perks = input.Perks

	UpdatedCampaign, err := s.repository.Update(campaign)
	if err!= nil{
		return campaign, errors.New("there something wrong")
	}
	return UpdatedCampaign,nil	




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

func (s *service) UploadCampaignImagePrimary(input CampaignImage) (Campaign, error){


	return Campaign{},errors.New("t")
}