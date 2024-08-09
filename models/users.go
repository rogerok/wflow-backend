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
	CreatedAt   time.Time  `json:"-" db:"created_at"`
	Email       string     `json:"email" db:"email"`
	FirstName   string     `json:"firstName" db:"first_name"`
	ID          uuid.UUID  `json:"id" db:"id"`
	LastName    *string    `json:"lastName" db:"last_name"`
	MiddleName  *string    `json:"middleName" db:"middle_name"`
	Pseudonym   Pseudonym  `json:"pseudonym" db:"pseudonym"`
	SocialLinks Social     `json:"socialLinks" db:"socialLinks"`
	UpdatedAt   time.Time  `json:"-" db:"updated_at"`
	BornDate    *time.Time `json:"bornDate" db:"born_date"`
}

func (s *Social) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &s)
}

func (p *Pseudonym) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &p)
}

func (p *Pseudonym) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (s *Social) Value() (driver.Value, error) {
	return json.Marshal(s)
}
