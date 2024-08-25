package model

// import "time"

type User struct {
	ID          string `json:"id" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Email       string `json:"email,omitempty" binding:"required,email"`
	Password    string `json:"password,omitempty" binding:"required,alphanum"`
	FirstName   string `json:"firstName,omitempty" binding:"required,alpha"`
	LastName    string `json:"lastName,omitempty" binding:"required,alpha"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"required,numeric"`
	// Wallet           Wallet    `gorm:"references:ID" json:"wallet"`
	// RegistrationDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"registrationDate" binding:"omitempty"`
	// ProfilePicture   []byte    `json:"profilePicture" binding:"omitempty"`
	// LastLogin        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastLogin" binding:"omitempty"`
}
