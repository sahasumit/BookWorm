package main

import (
	"controller"
	"fmt"
	"html/template"
	"model/dbcon"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
//var db *database.Db

var err error
var currentloggedin int //id of currently logged in user

//Member Home page
func Member(res http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("HTMLS/member/member.html")
	t.Execute(res, nil)
	//	UnsubscripeBook(2)

}

//publisher home page
func Publisher(res http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("HTMLS/publisher/publisher.html")
	t.Execute(res, nil)

}

//calling contact us

//html page handler
func HtmlHandler() {

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads")))) //file server for raw file serving inside html

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/signup", controller.SignUp)
	http.HandleFunc("/about", controller.About)
	http.HandleFunc("/contact", controller.Contact)
	http.HandleFunc("/member", controller.Member)
	http.HandleFunc("/publisher", controller.Publisher)
	http.HandleFunc("/my-unpublished-book", controller.MyUnpublishedBook)
	http.HandleFunc("/admin", controller.Admin)
	http.HandleFunc("/publish-new-book", controller.PublishNewBook)
	http.HandleFunc("/my-published-book", controller.MyPublishedBook)
	http.HandleFunc("/user-list", controller.UserList)
	http.HandleFunc("/unpublishedbook", controller.UnpublishedBook)
	http.HandleFunc("/publishedbook", controller.PublishedBook)
	http.HandleFunc("/adminreviewbook", controller.AdminReviewBook)
	http.HandleFunc("/approve-book", controller.ApproveBook)
	http.HandleFunc("/reject", controller.RejectBook)
	http.HandleFunc("/update-book", controller.UpdateBook)
	http.HandleFunc("/view-book", controller.ViewBook)
	//http.HandleFunc("/subscribe-book", SubscribeBook)
	http.HandleFunc("/unpublish-book", controller.UnpublishBook)
	/*http.HandleFunc("/view-user", ViewUser)
	http.HandleFunc("/block-user", BlockUser)
	http.HandleFunc("/send-notification", SendNotification)
	http.HandleFunc("/submit-notification", SubmitNotification)
	*/
}

//datbase conecting

func main() {
	//test()

	controller.Pr()

	fmt.Println("Server runing at port 8080")
	//DbConnection() //connecting with database
	dbcon.DbConnection()

	HtmlHandler()
	//creating server
	http.ListenAndServe(":8080", nil)
}
