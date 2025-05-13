package models

import (
	"errors"
	"go/by/example/restful/api/db"
	"go/by/example/restful/api/utils"
)

type Users struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := ` 
    INSERT INTO users(email, password)
    VALUES (?,?)
  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(&u.Email, &hashedPassword)
	return err
}

func (u *Users) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("credentials failed")
	}

	complience := utils.CheckCompliance(retrivedPassword, u.Password)
	if !complience {
		return errors.New("credentials failed")
	}

	return nil
}
