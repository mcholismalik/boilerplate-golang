# boilerplate-golang
Boilerplate Golang with Clean Architecture
- HOST : http://localhost:3030

## Pre Requisite
- Go version 1.19

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
This project built in clean architecture that contains some layer :
1. Driver   
2. Factory 
3. Delivery
4. Repository
5. Usecase
6. Model

# Packages
This project have some existing driver :
1. Http (rest, ws, web)
2. Database (postgres, mysql)
3. Elasticsearch
4. Firebase
5. Sentry
6. Websocket
7. Cron 

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
3. Webocket
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
