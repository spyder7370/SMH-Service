package models

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	StackTrace  string `json:"stackTrace"`
}
