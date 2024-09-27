package models

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Pseudonym struct {
	FirstName *string `json:"firstName" db:"pseudonym_first_name"`
	LastName  *string `json:"lastName" db:"pseudonym_last_name"`
}

type Social struct {
	Instagram *string `json:"instagram" db:"social_instagram"`
	Telegram  *string `json:"telegram" db:"social_telegram"`
	TikTok    *string `json:"tiktok" db:"social_tiktok"`
	Vk        *string `json:"vk" db:"social_vk"`
}

type User struct {
	BornDate    *time.Time `json:"-" db:"born_date"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	Email       string     `json:"email" db:"email"`
	FirstName   string     `json:"firstName" db:"first_name"`
	Id          uuid.UUID  `json:"id" db:"id"`
	LastName    *string    `json:"lastName" db:"last_name"`
	MiddleName  *string    `json:"middleName" db:"middle_name"`
	Password    string     `json:"-" db:"password"`
	Pseudonym   Pseudonym  `json:"pseudonym" db:"pseudonym"`
	SocialLinks Social     `json:"socialLinks" db:"socialLinks"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
}

func (s *Social) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &s)
}

func (s *Social) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (p *Pseudonym) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &p)
}

func (p *Pseudonym) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type UserQueryParams struct {
	PaginationQuery `json:"-"`
	OrderBy         string `json:"-" default:"createdAt desc"`
}
