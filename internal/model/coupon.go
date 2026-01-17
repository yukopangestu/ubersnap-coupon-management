package model

type Coupon struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Code            string  `json:"code"`
	DiscountType    string  `json:"discount_type"`
	DiscountValue   float64 `json:"discount_value"`
	MinPurchase     float64 `json:"min_purchase"`
	MaxDiscount     float64 `json:"max_discount"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
	IsActive        bool    `json:"is_active"`
	Amount          float64 `json:"amount"`
	RemainingAmount float64 `json:"remaining_amount"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type GetSingleCouponDTO struct {
	Name            string   `json:"name"`
	Amount          float64  `json:"amount"`
	RemainingAmount float64  `json:"remaining_amount"`
	ClaimedBy       []string `json:"claimed_by"`
}
