package repository

import (
	"github.com/username/go-webapp/internal/model"
	"gorm.io/gorm"
)

// CouponRepository defines the interface for coupon data access.
type CouponRepository interface {
	GetByName(name string) (*model.Coupon, error)
}

// couponRepository implements CouponRepository.
type couponRepository struct {
	db *gorm.DB
}

// NewCouponRepository creates a new instance of couponRepository.
func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{db: db}
}

// GetByName retrieves a coupon by its name.
func (r *couponRepository) GetByName(name string) (*model.Coupon, error) {
	var coupon model.Coupon
	if err := r.db.Where("name = ?", name).First(&coupon).Error; err != nil {
		return nil, err
	}
	return &coupon, nil
}
