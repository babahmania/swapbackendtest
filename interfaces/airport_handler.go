package interfaces

import (
	"strconv"
	"swapbackendtest/application"
	"swapbackendtest/domain/entity"
	"swapbackendtest/infrastructure/auth"

	"github.com/gofiber/fiber/v2"
)

//Airports struct defines the dependencies that will be used
type Airports struct {
	ap application.AirportAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//Airports constructor
func NewAirports(ap application.AirportAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Airports {
	return &Airports{
		ap: ap,
		rd: rd,
		tk: tk,
	}
}

func (s *Airports) SaveAirport(c *fiber.Ctx) error {
	var airport entity.Airport

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
	newAirport, err := s.ap.SaveAirport(&airport)
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
		"data":    newAirport,
	})
}

func (s *Airports) GetAirports(c *fiber.Ctx) error {
	var Airports = entity.Airports{}
	var err error
	Airports, err = s.ap.GetAirports()
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
		"data":    Airports,
	})
}

func (s *Airports) GetAirport(c *fiber.Ctx) error {
	airportId, err := strconv.ParseUint(c.Params("airport_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	airport, err := s.ap.GetAirport(airportId)
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
