package model

type Country struct {
	Code string `json:"code" gorm:"primary_key:true"`
	Name string `json:"name"`
}
