package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)


type Service interface{
	GenerateToken(userID int) (string, error)
}

type jwtService struct{


}

func NewService() *jwtService{
	return &jwtService{}
}

var SECRET_KEY =[]byte(os.Getenv("SECRET_KEY"))


//var secret_key = (os.Getenv(SECRET_KEY))

func (h *jwtService) GenerateToken(userID int) (string, error){

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil{
		return signedToken, err
	}
	return signedToken, nil
}

func(h *jwtService)ValidateToken(encodedtoken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedtoken,func(t *jwt.Token) (interface{}, error) {
			_,ok :=t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("invalid token")
			}
			
			return SECRET_KEY,nil


		})
		if err!= nil{
			return token,err

		}
		return token , nil
	}