package interfaces

import (
	"strconv"
	"swapbackendtest/application"
	"swapbackendtest/domain/entity"
	"swapbackendtest/infrastructure/auth"

	"github.com/gofiber/fiber/v2"
)

//Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//Users constructor
func NewUsers(us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

func (s *Users) SaveUser(c *fiber.Ctx) error {
	var user entity.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": validateErr,
		})
	}
	newUser, err := s.us.SaveUser(&user)
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
		"data":    newUser.PublicUser(),
	})
}

func (s *Users) GetUsers(c *fiber.Ctx) error {
	var users = entity.Users{}
	var err error
	users, err = s.us.GetUsers()
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
		"data":    users.PublicUsers(),
	})
}

func (s *Users) GetUser(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("user_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"success": false,
			"message": err.Error(),
		})
	}
	user, err := s.us.GetUser(userId)
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
		"data":    user.PublicUser(),
	})
}
