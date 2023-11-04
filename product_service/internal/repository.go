package internal

import (
	"errors"
	"github.com/AminN77/upera_test/product_service/pkg/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

var (
	ErrAddProduct    = errors.New("could not add product")
	ErrFetchProduct  = errors.New("could not fetch product")
	ErrUpdateProduct = errors.New("could not update product")
)

// Repository interface is implementation of the famous repository pattern for decoupling the
// database from the domain
type Repository interface {
	Add(p *Product) (*Product, error)
	Update(up *Product, id int) (*Product, *Product, []string, error)
	Fetch(id int) (*Product, error)
}

type postgresRepository struct {
	conn *gorm.DB
}

func NewPostgresRepository() Repository {
	conn, err := postgres.NewPostgresDb()
	if err != nil {
		log.Fatalln(err)
	}

	if err := conn.AutoMigrate(&Product{}); err != nil {
		log.Fatalln(err)
	}

	return &postgresRepository{
		conn: conn,
	}
}

func (pr *postgresRepository) Add(p *Product) (*Product, error) {
	p.Token, _ = uuid.NewUUID()
	if err := pr.conn.Create(p).Error; err != nil {
		log.Println("some database error occurred during add product, err:", err.Error())
		return nil, ErrAddProduct
	}
	return p, nil
}

func (pr *postgresRepository) Update(up *Product, id int) (*Product, *Product, []string, error) {
	var existingProduct *Product
	var beforeUpdate Product
	var changes []string

	if err := pr.conn.First(&existingProduct, id).Error; err != nil {
		log.Println("some database error occurred during fetch product, err:", err.Error())
		return nil, nil, changes, ErrUpdateProduct
	}

	if up.Token != existingProduct.Token {
		log.Println("the token is not correct for the product")
		return nil, nil, changes, ErrUpdateProduct
	}

	beforeUpdate = *existingProduct

	if up.Name != "" && up.Name != existingProduct.Name {
		existingProduct.Name = up.Name
		changes = append(changes, "Name")
	}

	if up.Description != "" && up.Description != existingProduct.Description {
		existingProduct.Description = up.Description
		changes = append(changes, "Description")
	}

	if up.Price != 0 && up.Price != existingProduct.Price {
		existingProduct.Price = up.Price
		changes = append(changes, "Price")
	}

	if up.Color != "" && up.Color != existingProduct.Color {
		existingProduct.Color = up.Color
		changes = append(changes, "Color")
	}

	if up.ImageUrl != "" && up.ImageUrl != existingProduct.ImageUrl {
		existingProduct.ImageUrl = up.ImageUrl
		changes = append(changes, "ImageUrl")
	}

	existingProduct.Token, _ = uuid.NewUUID()

	if err := pr.conn.Save(existingProduct).Error; err != nil {
		log.Println("some database error occurred during update product, err:", err.Error())
		return nil, nil, changes, ErrUpdateProduct
	}

	return &beforeUpdate, existingProduct, changes, nil
}

func (pr *postgresRepository) Fetch(id int) (*Product, error) {
	var product Product
	if err := pr.conn.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("some database error occurred during fetch product, err:", err.Error())
		return nil, ErrFetchProduct
	}

	return &product, nil
}
