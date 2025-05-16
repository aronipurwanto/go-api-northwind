package controllers

import (
	"github.com/aronipurwanto/go-api-northwind/internal/soap"
	"github.com/aronipurwanto/go-api-northwind/utils"
	"github.com/gofiber/fiber/v2"
)

type SOAPController struct {
	soapClient *soap.SOAPClient
}

func NewSOAPController(service *soap.SOAPClient) *SOAPController {
	return &SOAPController{soapClient: service}
}

// GetNumberInWords godoc
// @Summary Convert number to words
// @Tags SOAP
// @Produce json
// @Param number path string true "Number to convert"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.StandardErrorResponse
// @Router /soap/number-to-words/{number} [get]
func (c *SOAPController) GetNumberInWords(ctx *fiber.Ctx) error {
	number := ctx.Params("number")
	if number == "" {
		return utils.ErrorResponse(ctx, 400, "Number parameter is required", nil)
	}

	words, err := c.soapClient.CallNumberToWords(number)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to convert number", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 200, "Number converted", fiber.Map{
		"words": words,
	})
}
