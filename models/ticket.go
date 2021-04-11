package models

import "gorm.io/gorm"

// Ticket estructura que define el modelo ticket
// NOTA: gorm.Model otorga los campos ID, CreatedAt, UpdatedAt
type Ticket struct {
	gorm.Model
	Status string
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

// CreateTicket inserta los valores de un nuevo ticket en la base de datos
// en caso de fallo retorna una interface de error
func CreateTicket(db *gorm.DB, Ticket *Ticket) (err error) {
	// fmt.Println(Ticket)
	err = db.Create(Ticket).Error
	if err != nil {
		return err
	}
	db.Preload("User").Find(&Ticket)
	return nil
}

// GetTickets busca (todos) los tickets de la base de datos
// en caso de fallo retorna una interface de error
func GetTickets(db *gorm.DB, Ticket *[]Ticket) (err error) {
	err = db.Preload("User").Find(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// GetTicket busca un ticket por su respectivo id
// en caso de fallo retorna una interface de error
func GetTicket(db *gorm.DB, Ticket *Ticket, id string) (err error) {
	err = db.Where("id = ?", id).First(Ticket).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateTicket actualiza un ticket dado su id
func UpdateTicket(db *gorm.DB, Ticket *Ticket) (err error) {
	db.Save(Ticket)
	db.Preload("User").Find(&Ticket)
	return nil
}

// DeleteTicket elimina un ticket dado su id
func DeleteTicket(db *gorm.DB, Ticket *Ticket, id string) (err error) {
	db.Where("id = ?", id).Delete(Ticket)
	return nil
}
