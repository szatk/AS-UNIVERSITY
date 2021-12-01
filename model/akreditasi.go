package model

type Akreditasi struct {
	ID        uint    `json:"id" form:"id"`
	Akreditas string `json:"akreditas" form:"akreditas"`
}
