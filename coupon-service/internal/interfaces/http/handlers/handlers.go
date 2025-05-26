package handlers

// Handlers struct to store the customer handler and other handlers
type Handlers struct {
	CouponGenerateHandler *CouponGenerateHandler
}

// NewHandlers struct
func NewHandlers(couponGenerateHandler *CouponGenerateHandler) *Handlers {
	return &Handlers{
		CouponGenerateHandler: couponGenerateHandler,
	}
}
