package controller

import (
	"html/template"
	"log"
	"model"
	"net/http"
	"strconv"
)

//Admin Home page
func Admin(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/admin/admin.html")
	t.Execute(res, nil)
}

//list of  All unpublished book for admin
func UnpublishedBook(res http.ResponseWriter, req *http.Request) {
	var BL model.BookList
	//finding unpublished book id from 	database
	BL.Blist = model.GetBookList(0, 0)
	t, _ := template.ParseFiles("HTMLS/admin/unpublishedbook.html")
	t.Execute(res, BL)
}

//admin reviewing single book for publishing
func AdminReviewBook(res http.ResponseWriter, req *http.Request) {

	var book_id = req.URL.Query().Get("book")
	bid, _ := strconv.Atoi(book_id)
	book := model.GetBook(bid)
	log.Println("Admin checking book id : ", book_id)

	t, _ := template.ParseFiles("HTMLS/admin/adminreviewbook.html")
	t.Execute(res, book)
}

func ApproveBook(res http.ResponseWriter, req *http.Request) {
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be approved is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 1)
	http.Redirect(res, req, "/unpublishedbook", 301)
}

func RejectBook(res http.ResponseWriter, req *http.Request) {
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be rejected is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 2)
	http.Redirect(res, req, "/unpublishedbook", 301)
}
func UnpublishBook(res http.ResponseWriter, req *http.Request) {
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be unpublished is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 0)
	http.Redirect(res, req, "/publishedbook", 301)
}

func UserList(res http.ResponseWriter, req *http.Request) {

	type UserList struct {
		Ulist   []model.User
		Message string
	}

	var UL UserList
	UL.Ulist = model.GetUserList()
	t, _ := template.ParseFiles("HTMLS/admin/userlist.html")
	t.Execute(res, UL)

}
