package models

type Quotes struct {
	Id   string `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}
