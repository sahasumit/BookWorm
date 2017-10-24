# BookWorm
a platform where readers can read books and publishers can publish their books

![Linux Build Status](https://img.shields.io/travis/jekyll/jekyll/master.svg?label=Linux%20build)



This is a web app made with GO


## Installation
* Install GO language on your Linux-based machine form [here](https://golang.org/) and set your GOPATH
* From terminal enter `go get github.com/sahasumit/BookWorm`
* Now enter `cd $GOPATH/src/github.com/sahasumit/BookWorm`
* Edit the `config.json` file with your `host`, `port`, `database.host`, `database.port` `database.schema` `database.username`, `database.password`
* Now run the `ALL THE .sql files from "SQL-DUMP" folder` in you MySQL Workbench to create the full database
* Finally From terminal enter `go run main.go` and go to browser
-----

 ### Admin login credentials are already in use from given database, \\will be updated
  #### Admin1
 * Email: `sahasumit@gmail.com`
 * Password: `adminsumit`
 #### Admin2
* Email: `admin2@gmail.com`
* Password: `admin2`


 ### Other Libraries used

 * `github.com/gorilla/securecookie`
 * `github.com/gorilla/mux`
 * `github.com/go-sql-driver/mysql`
 -------
