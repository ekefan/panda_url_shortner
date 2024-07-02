package database

import (
	"fmt"

	"gorm.io/gorm"
)

// CreateUserArgs args for  creating a new user
type CreateUserArgs struct {
	Name string `json:"name"`
	Email string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

// CreateUser database query function for creating a new user row
// on success it returns a USER object and a nil error
func (s *Query) CreateUser(args CreateUserArgs) (USER, error){
	userRow := USER{Name: args.Name, Email: args.Email, Password: args.HashedPassword}
	result := s.db.Create(&userRow)
	return userRow, result.Error
}

// GetUserArgs args for getting user from the database
type GetUserArgs struct {
	Name string `json:"name"`
}

// GetUser gets a USER object from the database
// on sucess it returns the USER object specified by the GetUserArgs
func (s *Query) GetUser(args GetUserArgs) (USER, error){
	userRow := USER{}
	result := s.db.Where("name = ?", args.Name).First(&userRow)
	return userRow, result.Error
}

// TxUserArgs, fields need to update database user.name
type TxUserArgs struct {
	UserID uint `json:"user_id"`
	UserName string `json:"name"`
	NameUpdate string `json:"name_update"`
}

// TxUpdateUser database query to update username
func (s *Query) TxUpdateUser(args TxUserArgs) (USER, error) {
	userRow := USER{}
	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		//make query to update user.name in the database
		
		result := tx.Model(&userRow).Where("id = ?", args.UserID).Update("name", args.UserName)

		// handle error if any.
		if result.Error != nil {
			return result.Error
		}

		//check if there is a change affected else the transaction did not occur
		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows afffected, possible invalid user id: %v", args.UserID)
		}

		// update the url name from the url table
		result = tx.Model(&URL{}).Where("owner = ?", args.UserName).Update("owner", userRow.Name)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	// handle transaction errors
	if txErr != nil {
		return userRow, txErr
	}
	arg := GetUserArgs{ 
		Name: args.NameUpdate,
	}
	userRow, err := s.GetUser(arg)
	if err != nil {
		return userRow, err
	}
	return userRow, err
}


func (s *Query)TxDeleteUser(args TxUserArgs) error {
	userRow := USER{}
	urlRow := URL{}
	txErr := s.db.Transaction(
		func(tx *gorm.DB) error {
			result := tx.First(&userRow, "id = ?", args.UserID)
			if result.Error != nil {
				return result.Error
			}
			result = tx.First(&urlRow, "onwer = ?", args.UserName)
			if result.Error != nil {
				return result.Error
			}
			// delete all urls of the user
			result = tx.Where("owner = ?", userRow.Name).Delete(urlRow)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("no rows affectd for all urs")
			}
	
			//delete the user
			result = tx.Where("id = ?", userRow.ID).Delete(userRow)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("no rows affectd for all urs")
			}
			return nil
		})

		if txErr != nil {
			return txErr
		}
		return nil
}