package controllers

import (
	"github.com/aronipurwanto/go-api-northwind/internal/soap"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type SOAPController struct {
	client *soap.SOAPClient
}

func NewSOAPController(client *soap.SOAPClient) *SOAPController {
	return &SOAPController{client: client}
}

// ConvertNumber godoc
// @Summary Convert number to words using SOAP
// @Tags SOAP
// @Accept json
// @Produce json
// @Param number path int true "Number"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /soap/number-to-words/{number} [get]
func (s *SOAPController) ConvertNumber(c *fiber.Ctx) error {
	numStr := c.Params("number")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number"})
	}

	result, err := s.client.NumberToWords(num)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"words": result})
}
