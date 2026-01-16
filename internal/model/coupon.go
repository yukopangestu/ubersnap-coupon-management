package model

type Coupon struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Code          string  `json:"code"`
	DiscountType  string  `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	MinPurchase   float64 `json:"min_purchase"`
	MaxDiscount   float64 `json:"max_discount"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	IsActive      bool    `json:"is_active"`
	Stock         float64 `json:"stock"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
