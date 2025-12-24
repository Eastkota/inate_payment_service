package model
import (
	"time"
	"github.com/google/uuid"
)

type ContextKey string

const (
    TokenKey   ContextKey = "token_id"
    UserKey    ContextKey = "user"
    RequestKey ContextKey = "http_request"
)

type User struct {
    ID             uuid.UUID `json:"id"`
    UserIdentifier string    `json:"user_identifier"`
    Email          string    `json:"email"`
    MobileNo       string    `json:"mobile_no"`
    Password       string    `json:"password_hash"`
    Status         string    `json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

type UserActivity struct {
    ID  uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    Activity string `gorm:"type:varchar" json:"activity"`
    UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
    Count   int     `json:"count" gorm:"type:integer"`
    Month           int         `json:"month" gorm:"type:integer"`             
    Year            int         `json:"year" gorm:"type:integer"`

    User    *User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}