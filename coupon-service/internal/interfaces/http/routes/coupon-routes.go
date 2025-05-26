package routes

import (
	"coupon/internal/interfaces/http/handlers"
	"net/http"
)

func NewChargebackRoutes(handlers *handlers.CouponGenerateHandler) []Route {
	return []Route{
		{
			URI:          "/v1/coupons-generate",
			Method:       http.MethodPost,
			Handler:      handlers.CouponGenerate,
			RequiresAuth: false,
		},
		// Add other routes as needed
	}
}
