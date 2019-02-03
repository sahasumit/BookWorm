package main

import (
	"log"
	"net/http"

	"github.com/sahasumit/BookWorm/Api"
	"github.com/sahasumit/BookWorm/Controller"
	"github.com/sahasumit/BookWorm/model/configs"
	"github.com/sahasumit/BookWorm/model/dbcon"
	"github.com/sahasumit/BookWorm/view"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var err error

func HtmlHandlerMux() {

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))    //file server for raw file serving inside html
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource")))) //file server for raw file serving inside html

	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/logout", controller.Logout)
	router.HandleFunc("/signup", controller.SignUp)
	router.HandleFunc("/about", controller.About)
	router.HandleFunc("/contact", controller.Contact)
	router.HandleFunc("/user-home", controller.UserHome)
	router.HandleFunc("/my-un-published-book", controller.MyUnPublishedBook)
	router.HandleFunc("/publish-new-book", controller.PublishNewBook)
	router.HandleFunc("/my-published-book", controller.MyPublishedBook)
	router.HandleFunc("/user-list", controller.UserList)
	router.HandleFunc("/un-published-book", controller.UnPublishedBook)
	router.HandleFunc("/publishedbook", controller.PublishedBook)
	router.HandleFunc("/admin-review-book", controller.AdminReviewBook)
	router.HandleFunc("/approve-book", controller.ApproveBook)
	router.HandleFunc("/reject", controller.RejectBook)
	router.HandleFunc("/update-book", controller.UpdateBook)
	router.HandleFunc("/view-book", controller.ViewBook)
	router.HandleFunc("/subscribe-book", controller.SubscribeBook)
	router.HandleFunc("/unsubscribe-book", controller.UnsubscribeBook)
	router.HandleFunc("/unpublish-book", controller.UnpublishBook)
	http.Handle("/", router)
	router.HandleFunc("/send-notification", controller.SendNotification)
	router.HandleFunc("/post-notification", controller.PostNotification)
	router.HandleFunc("/user-control", controller.UserControl)
	router.HandleFunc("/read-book", controller.ReadBook)
	router.HandleFunc("/Users", Api.GetUsers)
}

var router = mux.NewRouter()

func main() {

	Config := configs.LoadConfiguration("config.json")
	view.Init()
	//DbConnection() //connecting with database
	dbcon.DbConnection(Config)
	//test()

	HtmlHandlerMux()
	//creating server

	log.Println("sumit saha  shawon rrr Server runing! go to", Config.Server.Host+":"+Config.Server.Port)
	err := http.ListenAndServe(Config.Server.Host+":"+Config.Server.Port, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("After Starting Server")
}
