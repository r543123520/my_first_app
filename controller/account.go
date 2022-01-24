package controller

import (
	"first_app/database"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	ID   int
	Name string `json:"name" validate:"required"`
}

func init() {
	database.DB.AutoMigrate(&Account{})
}

func CreateAccount(account Account) (Account, error) {
	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		return account, err
	}

	result := database.DB.Create(&account)

	return account, result.Error
}

func UpdateAccount(account Account, account_id int) (Account, error) {
	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		return account, err
	}

	account.ID = account_id

	result := database.DB.Save(&account)

	return account, result.Error
}

func GetAccounts() []Account {
	var result []Account
	database.DB.Raw("SELECT * FROM accounts").Scan(&result)

	return result
}

func DeleteAccount(account_id int) error {
	account := Account{
		ID: account_id,
	}
	result := database.DB.Delete(account)

	return result.Error
}
