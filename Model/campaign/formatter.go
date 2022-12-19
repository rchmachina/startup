package campaign



type CampaignFormatter struct{
	ID					int 	`json:"id" `
	UserID 				int	`json:"userid"`
	Name				string	`json:"name"`
	ShortDescription	string	`json:"short_description"`
	ImageUrl			string	`json:"image_url"`
	GoalAmount			int		`json:"goal_ammount"`
	CurrentAmount		int		`json:"current_ammount"`
	Slug				string	`json:"slug"`
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