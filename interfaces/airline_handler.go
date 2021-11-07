package interfaces

import (
	"strconv"
	"swapbackendtest/application"
	"swapbackendtest/domain/entity"
	"swapbackendtest/infrastructure/auth"

	"github.com/gofiber/fiber/v2"
)

//Airlines struct defines the dependencies that will be used
type Airlines struct {
	al application.AirlineAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//Airlines constructor
func NewAirlines(al application.AirlineAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Airlines {
	return &Airlines{
		al: al,
		rd: rd,
		tk: tk,
	}
}

func (s *Airlines) SaveAirline(c *fiber.Ctx) error {
	var airport entity.Airline
	err := c.BodyParser(&airport)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	//validate the request:
	validateErr := airport.Validate("")
	if len(validateErr) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": validateErr,
		})
	}
	newAirline, err := s.al.SaveAirline(&airport)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"success": true,
		"data":    newAirline,
	})
}

func (s *Airlines) GetAirlines(c *fiber.Ctx) error {
	var Airlines = entity.Airlines{}
	var err error
	Airlines, err = s.al.GetAirlines()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"success": true,
		"data":    Airlines,
	})
}

func (s *Airlines) GetAirline(c *fiber.Ctx) error {
	airportId, err := strconv.ParseUint(c.Params("airline_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	airport, err := s.al.GetAirline(airportId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"success": true,
		"data":    airport,
	})
}
