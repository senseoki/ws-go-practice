package dblayer

import "github.com/senseoki/ws-go-practice/gomusic/backend/models"

type DBLapper interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutuserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
