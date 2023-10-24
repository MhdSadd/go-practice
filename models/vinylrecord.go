package models


type Vinylrecord struct {
	Id int `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Image string `json:"image"`
}

