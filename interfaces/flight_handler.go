package interfaces

import (
	"fmt"
	"strconv"
	"swapbackendtest/application"
	"swapbackendtest/domain/entity"
	"swapbackendtest/infrastructure/auth"

	"github.com/gofiber/fiber/v2"
)

//Flights struct defines the dependencies that will be used
type Flights struct {
	fl      application.FlightAppInterface
	userApp application.UserAppInterface
	rd      auth.AuthInterface
	tk      auth.TokenInterface
}

//Flights constructor
func NewFlights(fl application.FlightAppInterface, uApp application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Flights {
	return &Flights{
		fl:      fl,
		userApp: uApp,
		rd:      rd,
		tk:      tk,
	}
}

func (s *Flights) SaveFlight(c *fiber.Ctx) error {
	//check is the user is authenticated first
	metadata, err := s.tk.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"success": false,
			"message": "unauthorized",
		})
	}
	//lookup the metadata in redis:
	userId, err := s.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"success": false,
			"message": "unauthorized",
		})
	}
	//check if the user exist
	_, err = s.userApp.GetUser(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": "user not found, unauthorized",
		})
	}

	var flight entity.Flight

	err = c.BodyParser(&flight)
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	//validate the request:
	validateErr := flight.Validate("")
	if len(validateErr) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": validateErr,
		})
	}
	flight.UserIDSubmit = userId
	newFlight, err := s.fl.SaveFlight(&flight)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"success": true,
		"data":    newFlight,
	})
}

func (s *Flights) GetFlights(c *fiber.Ctx) error {
	var Flights = entity.Flights{}
	var err error
	Flights, err = s.fl.GetFlights()
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
		"data":    Flights,
	})
}

func (s *Flights) GetFlight(c *fiber.Ctx) error {
	flightId, err := strconv.ParseUint(c.Params("flight_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	flight, err := s.fl.GetFlight(flightId)
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
		"data":    flight,
	})
}
func (s *Flights) GetFlightLocation(c *fiber.Ctx) error {
	var Flights = entity.Flights{}
	var err error
	originId, err := strconv.ParseUint(c.Params("origin_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	destinationId, err := strconv.ParseUint(c.Params("destination_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	qtyTransit, err := strconv.Atoi(c.Params("qty_transit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	Flights, err = s.fl.GetFlightLocation(originId, destinationId, qtyTransit)
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
		"data":    Flights,
	})
}
func (s *Flights) GetFlightLocationFind(c *fiber.Ctx) error {
	var Flights = entity.Flights{}
	var err error
	originId, err := strconv.ParseUint(c.Params("origin_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	destinationId, err := strconv.ParseUint(c.Params("destination_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	qtyTransit, err := strconv.Atoi(c.Params("qty_transit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	Flights, err = s.fl.GetFlightLocationFind(originId, destinationId, qtyTransit)
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
		"data":    Flights,
	})
}
