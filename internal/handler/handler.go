package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/username/go-webapp/internal/service"
)

type CouponHandler struct {
	service service.CouponService
}

func NewCouponHandler(service service.CouponService) *CouponHandler {
	return &CouponHandler{service: service}
}

func (h *CouponHandler) GetCouponDetail(c echo.Context) error {
	name := c.Param("name")
	coupon, err := h.service.GetCoupon(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
	}
	return c.JSON(http.StatusOK, coupon)
}
