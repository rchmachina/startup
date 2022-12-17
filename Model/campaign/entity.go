package campaign


import "time"

type Campaign struct{
	ID					int64 	//`json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	UserID 				int64
	Name				string
	ShortDescription	string
	Description			string
	Perks				string
	BackerCount 		int
	GoalAmount			int
	Slug				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
}

type CampaignImage struct{
	ID 				int
	CampaingImage 	int
	FileName 		string
	Isprimary		string
	CreatedAt		time.Time
	UpdatedAt		time.Time




}

