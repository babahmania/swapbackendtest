package main

import (
	"log"
	"os"
	"strconv"
	"swapbackendtest/infrastructure/auth"
	"swapbackendtest/infrastructure/persistence"
	"swapbackendtest/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	/*
		//issue mysql can not be connect in docker
		//test connection use github.com/go-sql-driver/mysql
		dbDocker, err := sql.Open("mysql", "swap:Sw4p-1443@tcp(db:3306)/flight-app-fiber")
		// if there is an error opening the connection, handle it
		if err != nil {
			log.Print(err.Error())
		}
		fmt.Println("connect")
		defer dbDocker.Close()
		panic(err)
	*/

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	//services.Automigrate()

	redisService, err := auth.NewRedisDB(redis_host, redis_port, redis_password)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()

	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	airports := interfaces.NewAirports(services.Airport, redisService.Auth, tk)
	airlines := interfaces.NewAirlines(services.Airline, redisService.Auth, tk)
	flights := interfaces.NewFlights(services.Flight, services.User, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	//user routes
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.SendString("Flight Search Services, ok👋!")
	})
	router := app.Group("/api/v1/users", logger.New())
	router.Post("/", users.SaveUser)
	router.Get("/", users.GetUsers)
	router.Get("/:user_id", users.GetUser)

	//authentication routes
	router.Post("/login", authenticate.Login)
	router.Post("/logout", authenticate.Logout)
	router.Post("/refresh", authenticate.Refresh)

	//airport routes
	router = app.Group("/api/v1/airports", logger.New())
	router.Post("/", airports.SaveAirport)
	router.Get("/", airports.GetAirports)
	router.Get("/:airport_id", airports.GetAirport)

	//airline routes
	router = app.Group("/api/v1/airlines", logger.New())
	router.Post("/", airlines.SaveAirline)
	router.Get("/", airlines.GetAirlines)
	router.Get("/:airline_id", airlines.GetAirline)

	//flight routes
	router = app.Group("/api/v1/flights", logger.New())
	router.Post("/", flights.SaveFlight)
	router.Get("/", flights.GetFlights)
	router.Get("/:flight_id", flights.GetFlight)
	router.Get("/location/:origin_id/:destination_id/:qty_transit", flights.GetFlightLocation)

	router.Get("/search/:origin_id/:destination_id/:qty_transit/:date_flight/:sort_flight/:departure_time/:class_id/:airline_flight", func(c *fiber.Ctx) error {
		originId, err := strconv.ParseUint(c.Params("origin_id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"success": false,
				"message": "Origin / Departure ID is not valid",
			})
		}
		destinationId, err := strconv.ParseUint(c.Params("destination_id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"success": false,
				"message": "Destination ID is not valid",
			})
		}
		qtyTransit, err := strconv.Atoi(c.Params("qty_transit"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"success": false,
				"message": "Value Transit is not valid",
			})
		}
		dateFlight, err := time.Parse("2006-01-02", c.Params("date_flight"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"success": false,
				"message": "Flight Departure Date is not valid",
			})
		}
		sortFlight := c.Params("sort_flight")
		departureTime := c.Params("departure_time")
		classFlight := c.Params("class_id")

		airlineFlight, err := strconv.Atoi(c.Params("airline_flight"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"success": false,
				"message": "Airline ID is not valid",
			})
		}

		flights, err := services.GetFlightLocationFind(originId, destinationId, qtyTransit, dateFlight, sortFlight, departureTime, classFlight, airlineFlight)
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
			"data":    flights,
		})
	})

	//starting the application
	app_port := os.Getenv("API_PORT")
	if app_port == "" {
		app_port = "8888"
	}

	log.Fatal(app.Listen(":" + app_port))
}
