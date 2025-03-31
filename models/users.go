package models

import (
	"go/by/example/restful/api/db"
	"go/by/example/restful/api/utils"
)

type Users struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error{
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
