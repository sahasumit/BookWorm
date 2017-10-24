# BookWorm
a platform where readers can read books and publishers can publish their books

![Linux Build Status](https://img.shields.io/travis/jekyll/jekyll/master.svg?label=Linux%20build)



This is a web app made with GO!


## Installation 

* test in Linux based Operating System
* install GO form [Here](https://golang.org/)
* go to your `GOPATH` or workspace 
* clone this project `git clone https://github.com/user/repo.git'
* build packeges `dbcon`,`model`,`view`,`controller` using `go build` & `go install command`
* create database Schema Named `BookWorm` in MysQL and create database from given [sql dump](https://github.com/sahasumit/BookWorm/tree/master/Database/Dump20171024)
* set your database `address`, `username` and `password` from `src/model/dbcon/dbcon.go`
* run `go run main.go`


 #### default admin login credentials are, 
 * Email: `sahasumit288@gmail.com`
 * Password: `adminsumit`
 
 ### Other Libraby used to make this
 
 * `github.com/gorilla/securecookie` 
 * `github.com/gorilla/mux`
 * `github.com/go-sql-driver/mysql`












