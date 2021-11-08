Flight search system for swapbackendtest
# Features
- [X] Users as administrator and customer, CRUD, seacrh / booking
- [X] User Login use JWT Token
- [X] Airport or city for CRUD, select / search booking
- [X] Airline for select / search booking, CRUD
- [X] Flight for select / search booking, CRUD, only user login and as administrator can be insert new flight schedule
    - [X] Search FLight By :
        - Departure Date
        - Departure Location
        - Destination Location
        - Departure Time
        - Class Flight
        - Value Transit
        - Airline Name
    - [X] Order By Search :
        - Departure Time
        - Arrival Time
        - Airline Name
        - Price Ascending
        - Price Descending
- [ ] Aircraft management / crud
- [ ] City, Country, International Flight
- [ ] Order Booking
- [ ] Order Booking Cancel
- [ ] Swagger Documentation API
- [X] Automigrate Schema Database
- [X] Docker
- [X] Unit testing coverage
    - User Service :
        - TestSaveUser_Success : go test -timeout 30s -run ^TestSaveUser_Success$ swapbackendtest/infrastructure/persistence -v
        - TestSaveUser_Failure : go test -timeout 30s -run ^TestSaveUser_Failure$ swapbackendtest/infrastructure/persistence -v
        - TestGetUser_Success  : go test -timeout 30s -run ^TestGetUser_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetUserByEmailAndPassword_Success  : go test -timeout 30s -run ^TestGetUserByEmailAndPassword_Success$ swapbackendtest/infrastructure/persistence -v
        - All Test :  E:\Go\bin\go.exe test -timeout 30s -run ^(TestSaveUser_Success|TestSaveUser_Failure|TestGetUser_Success|TestGetUsers_Success|TestGetUserByEmailAndPassword_Success)$ swapbackendtest/infrastructure/persistence -v

    - Airline Service :
        - TestSaveAirline_Failure : go test -timeout 30s -run ^TestSaveAirline_Failure$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirline_Success : go test -timeout 30s -run ^TestGetAirline_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirlines_Success : go test -timeout 30s -run ^TestGetAirlines_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirlineByID_Success : go test -timeout 30s -run ^TestGetAirlineByID_Success$ swapbackendtest/infrastructure/persistence -v
        - All Test : E:\Go\bin\go.exe test -timeout 30s -run ^(TestSaveAirline_Failure|TestGetAirline_Success|TestGetAirlines_Success|TestGetAirlineByID_Success)$ swapbackendtest/infrastructure/persistence -v

    - Airport Service :
        - TestSaveAirport_Failure : go test -timeout 30s -run ^TestSaveAirport_Failure$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirport_Success : go test -timeout 30s -run ^TestGetAirport_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirport_Success : go test -timeout 30s -run ^TestGetAirports_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetAirportByID_Success : go test -timeout 30s -run ^TestGetAirportByID_Success$ swapbackendtest/infrastructure/persistence -v
        - All Test : E:\Go\bin\go.exe test -timeout 30s -run ^(TestSaveAirport_Failure|TestGetAirport_Success|TestGetAirports_Success|TestGetAirportByID_Success)$ swapbackendtest/infrastructure/persistence -v

    - Flight Service :
        - TestSaveFlight_Failure : go test -timeout 30s -run ^TestSaveFlight_Failure$ swapbackendtest/infrastructure/persistence -v
        - TestGetFlight_Success : go test -timeout 30s -run ^TestGetFlight_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetFlight_Success : go test -timeout 30s -run ^TestGetFlights_Success$ swapbackendtest/infrastructure/persistence -v
        - TestGetFlightByID_Success : go test -timeout 30s -run ^TestGetFlightByID_Success$ swapbackendtest/infrastructure/persistence -v
        - All Test : go test -timeout 30s -run ^(TestSaveFlight_Failure|TestGetFlight_Success|TestGetFlights_Success|TestGetFlightByID_Success)$ swapbackendtest/infrastructure/persistence -v

# Start 
go run main.go

# Docker
- docker compose build --no-cache
- docker compose up

Note / Issue : 
https://forum.golangbridge.org/t/golang-not-connecting-to-mysql-thought-jinzhu-gorm-and-hosted-by-docker/3162
MySQL does starts before Golang, but, Golang is faster about trying to connect to the DB,
in first time run docker compose up will be error,
restart conatiner docker for service swap_api and then service can be run.

if error is docker Error response from daemon: Ports are not available 50212
net stop winnat
net start winnat


# User Access 
User as admin
"email": "babahmania@gmail.com",
"password": "babahmania@gmail.com"

User as member
"email": "maniababah@gmail.com",
"password": "maniababah@gmail.com"

# Link for demo
Backend / endpoint
http://158.140.191.182:50212/api/v1/


Frontend / web
https://kareemlogic.com/karim_flight/
this web consume api from end point http://158.140.191.182:50212/api/v1/

# Documentation :
- Folder Name : docs
- schema      : fiber-app-flight-schema.pdf
- database    : flight-app-fiber-db.sql
- postman     : swapbackendtest.postman_collection.json

# Note :
- This server in my home, if internet down, this server / end point can not be access.
- Database auto migrate in docker
- Issue use gorm from github.com/jinzhu/gorm v1.9.12 change to gorm.io/gorm, can not be connect docker db mysql
    script code for Automigrate in last commit is not exist again.