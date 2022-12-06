package helper





type Response struct{
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}


type Meta struct{
	Message string `json:"message"`
	Code 	int    `json:"code"`
	Status 	string `json:"status"`
}
type Formater struct{
	Username string`json:"userName"`
	Token string `json:"Token"`
}

func Formatuser(username string, token string) Formater{
	formatuser:= Formater{
		Username: username,
		Token: token ,
	}
	return formatuser
}

func APIResponse(message string,code int,status string, data interface{})Response {
	meta := Meta{
		Message: message ,
		Code: code,
		Status : status,


	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
		
	}
	return jsonResponse

}

