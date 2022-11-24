package Models

type RegisterUserInput struct{
	Name		string 	`binding:"required"`
	Occupation	string	`binding:"required"`
	Email		string	`binding:"required"`
	Password	string	`binding:"required"`
	AvatarFileName	string	
	Token 		string
}