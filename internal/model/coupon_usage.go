package model

type CouponUsage struct {
	ID       string `json:"id"`
	CouponId string `json:"coupon_id"`
	UserID   string `json:"user_id"`
}
