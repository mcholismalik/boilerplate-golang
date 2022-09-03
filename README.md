# boilerplate-golang
Boilerplate Golang
- HOST : http://localhost:3030

## Pre Requisite
- Go version 1.18.3

## How To Run
``` bash
# mod
$ go mod tidy

# run 
$ ENV=DEV go run main.go

# open
$ Open url http://localhost:3030
```

## Architecture 
This project built in clean architecture that contains :
1. Factory   
2. Middleware 
3. Handler
4. Binder
5. Validation
6. Usecase
7. Repository
8. Model
9. Database
9. Migration
10. Seed

# Packages
This project have some existing packages :
1. Elasticsearch   
2. Firebase
3. Sentry
4. Ws
5. Cron
6. Util
7. Database (postgres, mysql)

# Examples
This project have some example for rest, ws, web :]

1. Rest
  - Auth 
    - Login
    - Register
  - Sample
    - Get (+ pagination, sort & filter)
    - GetByID
    - Create (+ transaction scope)
    - Update (+ transaction scope)
    - Delete
2. Web
  - Playground
  - Bubble
3. Ws
  - Course

## Documentation

Install environment
``` bash
# get swagger package 
$ go install github.com/swaggo/swag/cmd/swag@latest

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{url}}/swagger/index.html

# Author
Muhammad Cholis Malik