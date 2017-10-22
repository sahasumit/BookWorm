package main

import (
	"controller"
	"fmt"
	"html/template"
	"log"
	"model"
	"model/dbcon"
	"net/http"
	"view"

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

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))    //file server for raw file serving inside html
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource")))) //file server for raw file serving inside html
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/signup", controller.SignUp)
	http.HandleFunc("/about", controller.About)
	http.HandleFunc("/contact", controller.Contact)
	http.HandleFunc("/user-home", controller.UserHome)
	http.HandleFunc("/my-unpublished-book", controller.MyUnpublishedBook)
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
	http.HandleFunc("/subscribe-book", controller.SubscribeBook)
	http.HandleFunc("/unsubscribe-book", controller.UnsubscribeBook)
	http.HandleFunc("/unpublish-book", controller.UnpublishBook)
	/*http.HandleFunc("/view-user", ViewUser)//ar lagbe na view user
	/*http.HandleFunc("/block-user", BlockUser)
	http.HandleFunc("/send-notification", SendNotification)
	http.HandleFunc("/submit-notification", SubmitNotification)
	*/
}

//testing func
func test() {
	log.Println("Test method : ")
	var data model.Notification
	data.BookId = 6

	data.AdminNotification = "Your book is a choti book"
	model.SendNotification(data)
	//	model.SetActiveUser(6, 0)
}

//datbase conecting

func main() {
	//test()

	controller.Pr()
	view.Init()

	fmt.Println("Server runing at port 8080")
	//DbConnection() //connecting with database
	dbcon.DbConnection()
	test()
	HtmlHandler()

	//creating server
	http.ListenAndServe(":8080", nil)
}
