package models

type DevIsValidResp struct {
	IsValid bool `json:"isValid"`
}

type DevIsBindResp struct {
	BindCount uint `json:"bindCount"`
}

type JmemberResp struct {
	UserId string `json:"userId" form:"userId" ` //id
	Name   string `json:"name" form:"name" `     //昵称
	Phone  string `json:"phone" form:"phone" `   //号码
}
