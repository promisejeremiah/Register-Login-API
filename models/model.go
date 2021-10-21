package models

import "time"

type User struct {
	FirstName   string    `json:"firstname,omitempty" 	bson:"firstname"`
	LastName    string    `json:"lastname,omitempty"	bson:"lastname"`
	Email       string    `json:"email,omitempty" 		bson:"email"`
	Password    string    `json:"password,omitempty" 	bson:"password"`
	Token       string    `json:"token,omitempty"		bson:"token"`
	Fresh_token string    `json:"refresh_token,omitempty"	bson:"refresh_token"`
	Created_at  time.Time `json:"created_at,omitempty"	bson:"created_at"`
	Updated_at  time.Time `json:"updated_at,omitempty"	bson:"updated_at"`
	User_id     string    `json:"user_id,omitempty"		bson: "user_id"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Contact struct {
	Fullname string `json:"fullname,omitempty" 	bson:"fullname"`
	Email    string `json:"email,omitempty" 		bson:"email"`
	Message  string `json:"message,omitempty" 		bson:"message"`
}
