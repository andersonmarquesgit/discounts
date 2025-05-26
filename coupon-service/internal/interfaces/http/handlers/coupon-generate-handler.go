package handlers

import (
	"coupon/internal/interfaces/dto"
	"coupon/internal/presentation"
	"coupon/internal/usecases"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CouponGenerateHandler struct {
	couponGenerateUseCase *usecases.CouponGenerateUseCase
}

func NewCouponGenerateHandler(couponGenerateUseCase *usecases.CouponGenerateUseCase) *CouponGenerateHandler {
	return &CouponGenerateHandler{
		couponGenerateUseCase: couponGenerateUseCase,
	}
}

// CouponGenerate godoc
// @Summary Generate coupons for a campaign
// @Description Generate coupons for a campaign
// @Tags chargeback
// @Accept json
// @Produce json
// @Param chargeback body dto.CouponGenerateRequest true "Data of the coupon generate"
// @Success 202 {object} presentation.JSONResponse
// @Success 200 {object} presentation.JSONResponse
// @Failure 400 {object} presentation.JSONResponse
// @Failure 500 {object} presentation.JSONResponse
// @Router /v1/coupon-generate [post]
func (h *CouponGenerateHandler) CouponGenerate(w http.ResponseWriter, r *http.Request) {
	var req dto.CouponGenerateRequest

	// Decode the request body into the chargeback request struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		presentation.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Validate the chargeback request struct
	if err := validate.Struct(req); err != nil {
		presentation.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Get the request-trace-id from context
	traceID := r.Context().Value("request-trace-id").(string)

	// Coupon generate
	_, err := h.couponGenerateUseCase.CouponGenerate(req, traceID)
	if err != nil {
		presentation.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Return the chargeback send to processor
	presentation.WriteJSON(w, http.StatusOK, presentation.JSONResponse{
		Error:   false,
		Message: "Coupon Generate Success",
		Data:    nil,
	})
}
