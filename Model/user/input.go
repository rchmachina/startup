package Models



type RegisterUserInput struct {
	Name           string `json:"name" binding:"required" form:"name"`
	Occupation     string `json:"occupation" binding:"required" form:"occupation"`
	Email          string `json:"email" binding:"required" form:"email"`
	Password       string `json:"password" binding:"required" form:"password"`
	Token          string
}

type LoginInput struct {
	Email    string `json:"email" binding:"required" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}

type ChangeImageInput struct {
	ID    int `json:"id" binding:"required" form:"id"`
	Token string `json:"token" binding:"required" form:"token"`
	
}
type CheckEmailAvaible struct {
	Email string `json:"email" binding:"required"`
}