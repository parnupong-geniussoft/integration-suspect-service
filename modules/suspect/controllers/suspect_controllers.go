package controllers

import (
	"integration-suspect-service/modules/suspect/entities"
	"integration-suspect-service/modules/suspect/usecases"
	"integration-suspect-service/pkg/validators"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type suspectController struct {
	SuspectUsecase usecases.SuspectUsecase
	Validate       *validator.Validate
}

func NewSuspectController(r fiber.Router, suspectUsecase usecases.SuspectUsecase, validate *validator.Validate) {
	controllers := &suspectController{
		SuspectUsecase: suspectUsecase,
		Validate:       validate,
	}
	r.Post("/add_suspect", controllers.AddSuspect)
}

// @Summary Add a suspect
// @Description Add a suspect using the IngSuspectRequest
// @Tags Suspects
// @Accept json
// @Produce json
// @Param body body entities.IngSuspectRequest true "Suspect request object"
// @Success 200 {object} entities.SuspectResponse "OK"
// @Router /add_suspect [post]
func (h *suspectController) AddSuspect(ctx *fiber.Ctx) error {
	body := new(entities.IngSuspectRequest)

	outJson := &entities.SuspectResponse{
		Status:     entities.STATUS_SUSPECT_FAILED,
		StatusDesc: entities.STATUS_SUSPECT_FAILED_DESC,
		Errors:     "",
	}

	if err := ctx.BodyParser(body); err != nil {
		outJson.Errors = err.Error()
		return ctx.JSON(outJson)
	}

	validators.RegisterCustomValidators(h.Validate)
	if err := h.Validate.Struct(body); err != nil {
		outJson.Errors = err.Error()

		return ctx.JSON(outJson)
	}

	xCorrelationID := ctx.Get("x_correlation_id")
	resp, err := h.SuspectUsecase.SubmitKtbSuspect(body, xCorrelationID, ctx)
	if err != nil {
		x := err.Error()
		outJson.StatusDesc = "KTB error :" + x
		outJson.Errors = "KTB error :" + x
		return ctx.JSON(outJson)
	}

	return ctx.JSON(entities.SuspectResponse{
		Control:    resp.Control,
		Status:     resp.Status,
		StatusDesc: entities.STATUS_SUSPECT_SUCCESS_DESC,
	})
}
