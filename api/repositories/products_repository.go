package repositories

import (
	"go-rest-example/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductsRepository interface {
	Save(*models.Product) (*models.Product, error)
	FindById(uint64) (*models.Product, error)
	FindAll() ([]*models.Product, error)
	Update(*models.Product) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(metadata *Metadata) (*Pagination, error)
	Search(string) ([]*models.Product, error)
}

type productsRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepositoryImpl {
	return &productsRepositoryImpl{db}
}

func (r *productsRepositoryImpl) Save(product *models.Product) (*models.Product, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Create(product).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return product, tx.Commit().Error
}

func (r *productsRepositoryImpl) FindById(id uint64) (*models.Product, error) {
	product := &models.Product{}
	err := r.db.Debug().Model(&models.Product{}).Where("id = ?", id).Take(&product).Error
	return product, err
}

func (r *productsRepositoryImpl) FindAll() ([]*models.Product, error) {
	products := []*models.Product{}
	err := r.db.Debug().Model(&models.Product{}).Take(&products).Error
	return products, err
}

func (r *productsRepositoryImpl) Update(product *models.Product) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"name":        product.Name,
		"price":       product.Price,
		"quantity":    product.Quantity,
		"status":      product.Status,
		"category_id": product.CategoryID,
		"updated_at":  time.Now(),
	}
	err := tx.Debug().Model(&models.Product{}).Where("id = ?", product.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsRepositoryImpl) Delete(id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsRepositoryImpl) Count() (int64, error) {
	var count int64
	err := r.db.Debug().Model(&models.Product{}).Count(&count).Error
	return count, err
}

func (r *productsRepositoryImpl) Paginate(metadata *Metadata) (*Pagination, error) {
	products := []*models.Product{}

	err := r.db.Debug().
		Model(&models.Product{}).
		Offset(metadata.Offset).
		Limit(metadata.Limit).
		Find(&products).Error

	return &Pagination{
		Elements: products,
		Metadata: metadata,
	}, err
}

func (r *productsRepositoryImpl) Search(name string) ([]*models.Product, error) {
	products := []*models.Product{}

	err := r.db.Debug().
		Model(&models.Product{}).
		Where("name LIKE ?", "%"+name+"%").
		Find(&products).Error

	return products, err
}
