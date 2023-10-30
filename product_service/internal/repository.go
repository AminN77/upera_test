package internal

import (
	"errors"
	"github.com/AminN77/upera_test/product_service/pkg/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	ErrAddProduct    = errors.New("could not add product")
	ErrUpdateProduct = errors.New("could not update product")
)

// Repository interface is implementation of the famous repository pattern for decoupling the
// database from the domain
type Repository interface {
	Add(p *Product) error
	Update(p *Product) error
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

func (pr *postgresRepository) Add(p *Product) error {
	if err := pr.conn.Create(p).Error; err != nil {
		log.Println("some database error occurred during add product, err:", err)
		return ErrAddProduct
	}
	return nil
}

func (pr *postgresRepository) Update(up *Product) error {
	if err := pr.conn.Save(up).Error; err != nil {
		log.Println("some database error occurred during update product, err:", err)
		return ErrUpdateProduct
	}
	return nil
}
