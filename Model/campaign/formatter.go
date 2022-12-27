package Models

import (
	"fmt"
	"strings"
)


type CampaignFormatter struct{
	ID					int 	`json:"id" `
	UserID 				int	`json:"userid"`
	Name				string	`json:"name"`
	ShortDescription	string	`json:"short_description"`
	ImageUrl			string	`json:"image_url"`
	GoalAmount			int		`json:"goal_ammount"`
	CurrentAmount		int		`json:"current_ammount"`
	Slug				string	`json:"slug"`
	Perks[]				string `json:"perks"`
	
}




func FormatCompaign (campaign Campaign)(CampaignFormatter){
	
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.UserID =campaign.UserID
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.ImageUrl = ""
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.Slug = campaign.Slug
	
	var perkss []string
	for _, perk :=range strings.Split(campaign.Perks,","){
		perkss= append(perkss,perk)
	}
	
	CampaignFormatter.Perks = perkss

	if len(campaign.CampaignImage) >0 { 
		CampaignFormatter.ImageUrl = campaign.CampaignImage[0].FileName
	}
	return CampaignFormatter
}

func FormatCompaigns(Campaign []Campaign)([]CampaignFormatter){
	var campaignsFormatter []CampaignFormatter
	
	for _, campaign := range Campaign{
		campaignFormatter := FormatCompaign(campaign)
		campaignsFormatter =append(campaignsFormatter,campaignFormatter)
	}
	return campaignsFormatter
}




type CampaignFormatterDetail struct{
	ID					int 	`json:"id" `
	
	DetailUser		DetailUser 	`json:"detail_user"`
	Name				string	`json:"name_campaign"`
	ShortDescription	string	`json:"short_description"`
	ImageUrl	string	`json:"image_url"`
	SecondaryImageUrl[]	SecondaryImage	`json:"Secondary_image_url"`
	GoalAmount			int		`json:"goal_ammount"`
	CurrentAmount		int		`json:"current_ammount"`
	Slug				string	`json:"slug"`
	Perks				[]string	`json:"perk"`
	
}

type DetailUser struct{
	UserID 				int	`json:"userid"`
	UserName 			string	`json:"Username"`
	Avatar 				string 	`json:"Avatar"`
	
}



type SecondaryImage struct{
	Secondary	bool	`json:"is secondary"`
	ImageUrl	string	`json:"Secondary_image_url"`

}

func FormatCompaignSecondaryImage(images CampaignImage)(SecondaryImage){
	SecondaryImages :=SecondaryImage{}
	if images.IsPrimary !=1{
		SecondaryImages.ImageUrl = images.FileName
		SecondaryImages.Secondary = true
	}
	return SecondaryImages

}


func FormatCompaignDetail (campaign Campaign)(CampaignFormatterDetail){
	DetailUser := DetailUser{}
	DetailUser.Avatar = campaign.User.AvatarFileName
	DetailUser.UserName = campaign.User.Name
	DetailUser.UserID = campaign.UserID
	CampaignFormatterDetail := CampaignFormatterDetail{}
	CampaignFormatterDetail.ID = campaign.ID
	

	CampaignFormatterDetail.DetailUser = DetailUser
	CampaignFormatterDetail.Name = campaign.Name
	CampaignFormatterDetail.ShortDescription = campaign.ShortDescription
	
	CampaignFormatterDetail.GoalAmount = campaign.GoalAmount
	CampaignFormatterDetail.CurrentAmount = campaign.CurrentAmount
	CampaignFormatterDetail.Slug = campaign.Slug
	

	if len(campaign.CampaignImage) >0 { 
		CampaignFormatterDetail.ImageUrl = campaign.CampaignImage[0].FileName
		 for i:=1; i<len(campaign.CampaignImage);i++{
			SecondaryImage := FormatCompaignSecondaryImage(campaign.CampaignImage[i])
			CampaignFormatterDetail.SecondaryImageUrl  = append(CampaignFormatterDetail.SecondaryImageUrl , SecondaryImage )
		 }

	
	}
	var perkss []string
	for _, perk :=range strings.Split(campaign.Perks,","){
		perkss= append(perkss,perk)
	}
	fmt.Println(perkss)
	CampaignFormatterDetail.Perks = perkss

	
	return CampaignFormatterDetail
}


func FormatCompaignsDetail(Campaign []Campaign)([]CampaignFormatterDetail){
	var campaignsFormatterDetail []CampaignFormatterDetail
	
	for _, campaign := range Campaign{
		campaignFormatterDetails := FormatCompaignDetail(campaign)
		campaignsFormatterDetail =append(campaignsFormatterDetail,campaignFormatterDetails)
	}
	return campaignsFormatterDetail
}