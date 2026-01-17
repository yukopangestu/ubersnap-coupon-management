package service

import (
	"github.com/username/go-webapp/internal/model"
	"github.com/username/go-webapp/internal/repository"
)

// CouponService defines the interface for coupon business logic.
type CouponService interface {
	GetCoupon(name string) (*model.GetSingleCouponDTO, error)
}

// couponService implements CouponService.
type couponService struct {
	couponRepo      repository.CouponRepository
	couponUsageRepo repository.CouponUsageRepository
}

// NewCouponService creates a new instance of couponService.
func NewCouponService(couponRepo repository.CouponRepository, couponUsageRepo repository.CouponUsageRepository) CouponService {
	return &couponService{couponRepo: couponRepo, couponUsageRepo: couponUsageRepo}
}

// GetCoupon retrieves a coupon by name using the repository.
func (s *couponService) GetCoupon(name string) (*model.GetSingleCouponDTO, error) {
	var result model.GetSingleCouponDTO
	couponData, err := s.couponRepo.GetByName(name)

	if err != nil {
		return &result, err
	}

	// Fetch coupon usage to get claimed_by users
	couponUsages, err := s.couponUsageRepo.GetCouponUsagebyCouponId(couponData.ID)
	if err != nil {
		return nil, err
	}

	var claimedBy []string
	if couponUsages != nil {
		for _, usage := range *couponUsages {
			claimedBy = append(claimedBy, usage.UserID)
		}
	}

	result = model.GetSingleCouponDTO{
		Name:            couponData.Name,
		Amount:          couponData.Amount,
		RemainingAmount: couponData.RemainingAmount,
		ClaimedBy:       claimedBy,
	}

	return &result, nil
}
