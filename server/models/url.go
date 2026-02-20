package models

import(
	"time"
)

type Url struct {
	ID          int    		`json:"id"`
	UserID      string 		`json:"user_id"`
	ShortCode   string 		`json:"short_code"`
	OriginalUrl string 		`json:"original_url"`
	CreatedAt   time.Time 	`json:"created_at"`
	ExpiresAt   time.Time 	`json:"expires_at"`
	Clicks      int			`json:"clicks"`
}