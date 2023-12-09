# Go Starter 2024

A starter for Go API micro-services. 

## Pre-Requisites

You will need Docker and Go (1.7 or above) installed before setting up the app.

## Setup

You will need to install Go 1.7 or above. Then configure the following:

1. Change the configuration file

2. Set the ("ENVIRONMENT_NAME") environment variable, either using terminal or os.Setenv("ENVIRONMENT_NAME","dev").

3. Set the JWT secret env var ("JWT_SECRET")

4. In cmd/migration/main.go set up psn variable and then run it (go run main.go). It will create all tables, and necessary data, with a new account username/password admin/admin.

5. Run the app using:

```bash
go run cmd/api/main.go
```

## Installation

First, install PostgreSQL and InfluxDB.

```bash
$ docker pull postgres:15.1
$ docker run -d -p 5432:5432 --name postgresDB15 -e POSTGRES_PASSWORD=pass1234 postgres:15.1
$ docker start postgresDB15
```

## Running the App

```bash
go clean && go build && go run ./cmd/api/main.go
```

### Dependencies

```bash
|-------------------------------------|--------------------------------------------|--------------|
|             DEPENDENCY              |                  REPO URL                  |   LICENSE    |
|-------------------------------------|--------------------------------------------|--------------|
| github.com/labstack/echo/v4         | https://github.com/labstack/echo/v4        | MIT          |
| github.com/go-pg/pg                 | https://github.com/go-pg/pg                | bsd-2-clause |
| github.com/dgrijalva/jwt-go         | https://github.com/dgrijalva/jwt-go        | MIT          |
| github.com/rs/zerolog               | https://github.com/rs/zerolog              | MIT          |
| golang.org/x/crypto/bcrypt          | https://github.com/golang/crypto           |              |
| gopkg.in/yaml.v2                    | https://github.com/go-yaml/yaml            |              |
| gopkg.in/go-playground/validator.v8 | https://github.com/go-playground/validator | MIT          |
| github.com/lib/pq                   | https://github.com/lib/pq                  | Other        |
| github.com/nbutton23/zxcvbn-go      | https://github.com/nbutton23/zxcvbn-go     | MIT          |
| github.com/fortytw2/dockertest      | https://github.com/fortytw2/dockertest     | MIT          |
| github.com/stretchr/testify         | https://github.com/stretchr/testify        | Other        |
|-------------------------------------|--------------------------------------------|--------------|
```

1. Echo - HTTP 'framework'.
2. Go-Pg - PostgreSQL ORM
3. JWT-Go - JWT Authentication
4. Zerolog - Structured logging
5. Bcrypt - Password hashing
6. Yaml - Unmarshalling YAML config file
7. Validator - Request validation.
8. lib/pq - PostgreSQL driver
9. zxcvbn-go - Password strength checker
10. DockerTest - Testing database queries
11. Testify/Assert - Asserting test results

## Accessing Tools

Connecting to postgres:

```bash
$ docker exec -it -u postgres postgresDB15 psql
```

Connect to InfluxDB:

```bash
$ docker-compose --env-file dev_influxdb.env up
```

## Code Considerations

### Implementing CRUD of another table

Let's say you have a table named 'cars' that handles employee's cars. To implement CRUD on this table you need:

1. Inside `pkg/utl/model` create a new file named `car.go`. Inside put your entity (struct), and methods on the struct if you need them.

2. Create a new `car` folder in the (micro)service where your service will be located, most probably inside `api`. Inside create a file/service named car.go and test file for it (`car/car.go` and `car/car_test.go`). You can test your code without writing a single query by mocking the database logic inside /mock/mockdb folder. If you have complex queries interfering with other entities, you can create in this folder other files such as car_users.go or car_templates.go for example.

3. Inside car folder, create folders named `platform`, `transport` and `logging`.

4. Code for interacting with a platform like database (postgresql) should be placed under `car/platform/pgsql`. (`pkg/api/car/platform/pgsql/car.go`)

5. In `pkg/api/car/transport` create a new file named `http.go`. This is where your handlers are located. Under the same location create http_test.go to test your API.

6. In logging directory create a file named `car.go` and copy the logic from another service. This serves as request/response logging.

6. In `pkg/api/api.go` wire up all the logic, by instantiating car service, passing it to the logging and transport service afterwards.

### Implementing other platforms

Similarly to implementing APIs relying only on a database, you can implement other platforms by:

1. In the service package, in car.go add interface that corresponds to the platform, for example, Indexer or Reporter.

2. Rest of the procedure is same, except that in `/platform` you would create a new folder for your platform, for example, `elastic`.

3. Once the new platform logic is implemented, create an instance of it in main.go (for example `elastic.Client`) and pass it as an argument to car service (`pkg/api/car/car.go`).

### Running database queries in transaction

To use a transaction, before interacting with db create a new transaction:

```go
err := s.db.RunInTransaction(func (tx *pg.Tx) error{
    // Application service here
})
````

Instead of passing database client as `s.db` , inside this function pass it as `tx`. Handle the error accordingly.

## Project Structure

1. Root directory contains things not related to code directly, e.g. docker-compose, CI/CD, readme, bash scripts etc. It should also contain vendor folder, Gopkg.toml and Gopkg.lock if dep is being used.

2. Cmd package contains code for starting applications (main packages). The directory name for each application should match the name of the executable you want to have. An application may produce multiple binaries, therefore we use the Go convention of placing main package as a subdirectory of the cmd package. As an example, in a scheduler application the binary would be located under cmd/cron. It also loads the necessary configuration and passes it to the service initializers.

3. Rest of the code is located under /pkg. The pkg directory contains `utl` and 'microservice' directories.

4. Microservice directories, like api (naming corresponds to `cmd/` folder naming) contains multiple folders for each domain it interacts with, for example: user, car, appointment etc.

5. Domain directories, like user, contain all application/business logic and two additional directories: platform and transport.

6. Platform folder contains various packages that provide support for things like databases, authentication or even marshaling. Most of the packages located under platform are decoupled by using interfaces. Every platform has its own package, for example, postgres, elastic, redis, memcache etc.

7. Transport package contains HTTP handlers. The package receives the requests, marshals, validates then passes it to the corresponding service.

8. Utl directory contains helper packages and models. Packages such as mock, middleware, configuration, server are located here.

## Deployment

The following environment variables must be set:

* ENVIRONMENT_NAME
* PG_URL
* PG_USER
* PG_PASS
* V1_AUTH_USER
* V1_AUTH_PASS
* JWT_SECRET

## Contributing

We loosely follow a trunk-based protocol.

## Common Errors

## Notes
