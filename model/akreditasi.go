package model

type Akreditasi struct {
	ID        int    `json:"id" form:"id"`
	Akreditas string `json:"akreditas" form:"akreditas"`
}
