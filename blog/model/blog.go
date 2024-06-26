package model

type Blog struct{
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	Author string `json:"author"`
}