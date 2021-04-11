package models

import (
	"gorm.io/gorm"
)

// User estructura que define el modelo usurio
// NOTA: gorm.Model otorga los campos ID, CreatedAt, UpdatedAt
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser inserta los valores de un nuevo usuario en la base de datos
// en caso de fallo retorna una interface de error
func CreateUser(db *gorm.DB, User *User) (err error) {
	// fmt.Println(User)
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUsers busca (todos) los usuarios de la base de datos
// en caso de fallo retorna una interface de error
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser busca un usuario por su respectivo id
// en caso de fallo retorna una interface de error
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser actualiza un usuario dado su id
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// DeleteUser elimina un usuario dado su id
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
