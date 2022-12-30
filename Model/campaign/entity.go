package Models


import ("time"
		Models "campaign/Model/user")


type Campaign struct{
	ID					int 		`json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	UserID 				int			`json:"userid" binding:"required" form:"userid"`
	Name				string		`json:"name" gorm:"type:varchar(100)"`
	ShortDescription	string		`json:"short_description" gorm:"type:varchar(100)"`
	Description			string		`json:"long_description" gorm:"type:varchar(200)"`
	Perks				string		`json:"perks" gorm:"type:varchar(100)"`
	BackerCount 		int			`json:"Backer_count" `
	GoalAmount			int			`json:"goal_ammount"`
	CurrentAmount		int			`json:"current_ammount"`
	Slug				string		`json:"slug" gorm:"type:varchar(100)"`
	ImageUrl			string		`json:"image_url" gorm:"type:varchar(100)"`
	CreatedAt			time.Time 	
	UpdatedAt			time.Time 	
	CampaignImage		[]CampaignImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User				Models.User

	
}

type CampaignImage struct{
	ID 				int `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	CampaignID	 	int `json:"userid" binding:"required" form:"CampaignId"`
	FileName 		string	
	IsPrimary		int	`json:"primaryfile" form:"isprimary"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
}



