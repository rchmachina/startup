package Models

import (

	"time"
)


type User struct{
	ID				int64 	`json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name			string	`json:"name" gorm:"type:varchar(100)"`
	Occupation		string	`json:"occupation" gorm:"type:varchar(100)"`
	Email			string	`json:"email" gorm:"type:varchar(100);not_null"`
	PasswordHash	string	`json:"passwordhash" gorm:"type:varchar(100);not_null"`
	AvatarFileName	string	`json:"avatarfilename" gorm:"type:varchar(100);not_null"`
	Role 			string	`json:"rolename" gorm:"type:varchar(100);not_null"`
	Token 			string	`json:"token" gorm:"type:varchar(100);not_null"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	//campaign []campaign.Campaign
}

