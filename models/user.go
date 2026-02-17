package models

type User struct{
	Id int `json:"UserId"`
	Name string `json:"UserName"`
	Password string `json:"UserPassword"`
}
