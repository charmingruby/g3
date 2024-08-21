package v1

import (
	"github.com/charmingruby/g3/internal/common/api/api_rest"
	"github.com/charmingruby/g3/internal/common/custom_err"
	"github.com/charmingruby/g3/internal/telemetry/domain/dto"
	"github.com/charmingruby/g3/internal/telemetry/transport/rest/endpoint/v1/presenter"
	"github.com/gin-gonic/gin"
)

type CreateGPSRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (h *Handler) createGPSEndpoint(c *gin.Context) {
	var req CreateGPSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api_rest.NewPayloadError(c, err)
		return
	}

	dto := dto.CreateGPSInputDTO{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	output, err := h.telemetryService.CreateGPSUseCase(dto)
	if err != nil {
		validationErr, ok := err.(*custom_err.ErrValidation)
		if ok {
			api_rest.NewEntityError(c, validationErr)
			return
		}

		api_rest.NewInternalServerError(c, err)
		return
	}

	mappedGPS := presenter.DomainGPSToHTTP(output.GPS)

	api_rest.NewCreatedResponse(c, "gps", mappedGPS)
}
