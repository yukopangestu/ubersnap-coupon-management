package repository

import (
	"github.com/username/go-webapp/internal/model"
	"gorm.io/gorm"
)

type CouponUsageRepository interface {
	GetCouponUsagebyCouponId(couponId string) (*[]model.CouponUsage, error)
}

// couponRepository implements CouponRepository.
type couponUsageRepository struct {
	db *gorm.DB
}

// NewCouponUsageRepository creates a new instance of couponUsageRepository.
func NewCouponUsageRepository(db *gorm.DB) CouponUsageRepository {
	return &couponUsageRepository{db: db}
}

func (r *couponUsageRepository) GetCouponUsagebyCouponId(couponId string) (*[]model.CouponUsage, error) {
	var couponUsage []model.CouponUsage
	if err := r.db.Where("coupon_id = ?", couponId).Find(&couponUsage).Error; err != nil {
		return nil, err
	}

	return &couponUsage, nil
}
