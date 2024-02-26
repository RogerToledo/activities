# CRUD todo-api
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RogerToledo/todo-api)
![Static Badge](https://img.shields.io/badge/v.16-white?logo=postgresql&logoColor=white&label=Postgres&labelColor=blue&color=gray)


## Description
This is a simple CRUD API for a todo list. It was built using Golang and Postgres.
It was developed using a tutorial from [Aprenda Golang](https://www.youtube.com/watch?v=MD7b-iQMC24) as a base, and then I added some features and refactored some parts of the code.

## How to run
1. Clone this repository
2. Run compose-postgres or use your own Postgres database
    If you want to use your own Postgres, you must change database values on config.toml
3. Run ddl.sql to create the table    
2. Run `go run main.go` to start the server

## Author
[Roger Toledo](https://github.com/RogerToledo)