Flight search system for swapbackendtest
# Features
- [X] Users as administrator and customer, CRUD, seacrh / booking
- [X] User Login use JWT Token
- [X] Airport or city for CRUD, select / search booking
- [X] Airline for select / search booking, CRUD
- [X] Flight for select / search booking, CRUD, only user login and as administrator can be insert new flight schedule
    Search FLight By :
    - Departure Date
    - Departure Location
    - Destination Location
    - Departure Time
    - Class Flight
    - Value Transit
    - Airline Name
    Order By Search :
    - Departure Time
    - Arrival Time
    - Airline Name
    - Price Ascending
    - Price Descending
- [ ] Aircraft management / crud
- [ ] Order Booking
- [ ] Order Booking Cancel

# Start 
go run main.go

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
schema      : fiber-app-flight-schema.pdf
database    : fiber-app-flight-db.sql
postman     : swapbackendtest.postman_collection.json
# Note :
- This server in my home, if internet down, this server / end point can not be access.

- because form migrator can be create view view_flights schema, 
    create manual script use
        SELECT `flights`.`id`,
            `flights`.`flight_number`,
            `flights`.`airline_id`,
            `airlines`.`name` as airline_name,
            `airlines`.`image_name` as airline_image_name,
            `flights`.`origin_id`,
            origin_airport.`name` as origin_name,
            origin_airport.`code` as origin_code,
            `flights`.`destination_id`,
            destination_airport.`name` as destination_name,
            destination_airport.`code` as destination_code,
            `flights`.`aircraft_id`,
            `aircrafts`.`name` as aircraft_name,
            `flights`.`depart_datetime`,
            `flights`.`arrival_datetime`,
            `flights`.`duration`,
            `flights`.`price`,
            `flights`.`seats_available`,
            `flights`.`qty_transit`,
            `flights`.`flight_status`,
            `flights`.`user_id_submit`,
            `flights`.`user_id_update`,
            `flights`.`user_id_delete`,
            `flights`.`transit_first`,
            `flights`.`transit_second`,
            `flights`.`transit_third`,
            `flights`.`is_economy`,
            `flights`.`seats_available_economy`,
            `flights`.`is_premium_economy`,
            `flights`.`seats_available_premium_economy`,
            `flights`.`is_business`,
            `flights`.`seats_available_business`,
            `flights`.`is_first_class`,
            `flights`.`seats_available_first_class`,
            `flights`.`qty_baggage`,
            `flights`.`qty_cabin`,
            `flights`.`is_meal`,
            `flights`.`is_entertainment`,
            `flights`.`is_power_usb`,
            `flights`.`is_active`
        FROM `flight-app-fiber`.`flights`, `flight-app-fiber`.`airlines`, 
        `flight-app-fiber`.`airports` as origin_airport, `flight-app-fiber`.`airports` as destination_airport, 
        `flight-app-fiber`.`aircrafts` 
        where `flights`.`airline_id`=`airlines`.`id` and 
        `flights`.`origin_id`=origin_airport.`id` and
        `flights`.`destination_id`=destination_airport.`id` and 
        `flights`.`aircraft_id`=`aircrafts`.`id`
        order by `flights`.`depart_datetime`






