package Models

import (
	Models "campaign/Model/user"

)


type GetCampaignByID struct{
	ID	int `uri:"id" binding:"required" `

}

type CreateCampaign struct{
	Name				string		`json:"name" gorm:"type:varchar(100)" binding:"required" form:"name"`
	ShortDescription	string		`json:"short_description" gorm:"type:varchar(100)" binding:"required" form:"short_description"`
	Description			string		`json:"long_description" gorm:"type:varchar(200)" binding:"required"  form:"long_description"`
	Perks				string		`json:"perks" gorm:"type:varchar(100)" binding:"required"  form:"perks"`
	GoalAmount			int			`json:"goal_ammount" binding:"required"  form:"goal_ammount" gorm:"type:int(200)"`
	User				Models.User 
}

