package repository

import (
	"database/sql"
	"e-wallet-app/internal/model"
	// "fmt"
)

type UserRepository interface {
	BaseRepository[model.User]
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) Create(payload *model.User) error {
	_, err := r.db.Exec("INSERT INTO novel(id,username,email,password,firstName,lastName,phoneNumber)VALUES($1, $2, $3, $4, $5, $6, $7)", payload.ID, payload.Username, payload.Email, payload.Password, payload.FirstName, payload.LastName, payload.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
	// user := model.User{
	// 	ID:          payload.ID,
	// 	Username:    payload.Username,
	// 	Email:       payload.Email,
	// 	Password:    payload.Password,
	// 	FirstName:   payload.FirstName,
	// 	LastName:    payload.LastName,
	// PhoneNumber: payload.PhoneNumber,
	// Wallet: model.Wallet{
	// 	ID:      common.GenerateUUID(),
	// 	Name:    "sakupay",
	// 	Balance: 0,
	// },
	// }
	// if err := r.db.Create(&user).Error; err != nil {
	// 	return nil, err

	// }
	// fmt.Println(user)

	// return &user, nil
}
