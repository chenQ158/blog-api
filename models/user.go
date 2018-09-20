package models

type User struct {
	Id			int		`json:"ID"`
	Username	string	`json:"USERNAME"`
	Nickname	string	`json:"NICKNAME"`
}
