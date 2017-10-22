package controller

import (
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"model"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Pr() {
	fmt.Println("Hello from Package")
}

func Home(res http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("HTMLS/index.html")
	t.Execute(res, nil)
}

var Currentloggedin int

func Login(res http.ResponseWriter, req *http.Request) {
	log.Println("method login")
	if req.Method != "POST" {

		t, _ := template.ParseFiles("HTMLS/login.html")
		t.Execute(res, nil)
		return
	}

	//logging
	req.ParseForm()
	email := html.EscapeString(req.FormValue("email"))
	password := html.EscapeString(req.FormValue("password"))
	log.Println(time.Now().Format(time.RFC850), "User Login Attempt by: ", email, " ", password)
	var user model.User
	user = model.GetUser(email)

	if user.Email != email {
		log.Println("User not found")
		http.Redirect(res, req, "/login", 301)
		return
	}
	if user.Password != password {
		log.Println("Password does not match")
		http.Redirect(res, req, "/login", 301)
		return
	}

	//if user is block redirect him
	if user.IsActive == 0 {
		log.Println("User is blocked")
		http.Redirect(res, req, "/login", 301)
		return
	}

	Currentloggedin = user.UserId
	log.Println("Welcome success login id = " + strconv.Itoa(Currentloggedin))
	//redirect according to user type
	if user.UserType == "member" {
		http.Redirect(res, req, "/member", 301)
		return
	}

	if user.UserType == "publisher" {
		http.Redirect(res, req, "/publisher", 301)
		return
	}

	if user.UserType == "admin" {
		http.Redirect(res, req, "/admin", 301)
		return
	}
}

func SignUp(res http.ResponseWriter, req *http.Request) {
	//before clicking submit option
	if req.Method != "POST" {
		t, _ := template.ParseFiles("HTMLS/signup.html")
		t.Execute(res, nil)
		return
	}

	//getting signup information
	req.ParseForm()
	name := html.EscapeString(req.FormValue("name"))
	email := html.EscapeString(req.FormValue("email"))
	password1 := html.EscapeString(req.FormValue("password1"))
	password2 := html.EscapeString(req.FormValue("password2"))
	usertype := html.EscapeString(req.FormValue("UserType"))
	log.Println("Name ", name, "Email ", email, "password1 ", password1, "password2 ", password2, " Type ", usertype)

	//matching password for confirmation
	if password1 != password2 {
		println("Password does not match")
		http.Redirect(res, req, "/signup", 301)
		return
	}
	//checking mail used or not
	var emailexist string
	var user model.User
	user = model.GetUser(email)
	emailexist = user.Email
	if emailexist == email {
		println("Email already used")
		http.Redirect(res, req, "/signup", 301)
		return
	}
	//generating unique user id
	var user_id int
	user_id = model.GenerateID(1)
	user.Set(user_id, email, password1, name, 1, usertype)
	model.SetUser(user)
	println("Sign Up successfull ", user_id)
	//storing new user in database user tab
	println("Stored in database")
	http.Redirect(res, req, "/login", 301)
}

func Contact(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/contact.html")
	t.Execute(res, nil)
}
func About(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/about.html")
	t.Execute(res, nil)
}

func PublishedBook(res http.ResponseWriter, req *http.Request) {
	var BL model.BookList
	//finding unpublished book id from 	database
	BL.Blist = model.GetBookList(1, 0) // 1 - publishedbook, 0 - No specific user
	t, _ := template.ParseFiles("HTMLS/publishedbook.html")
	t.Execute(res, BL)
}

func MyPublishedBook(res http.ResponseWriter, req *http.Request) {
	var BL model.BookList
	log.Println("Method:MyPublishedBook -> Publisher id = ", Currentloggedin)
	//Take publisherid(Currentloggedin) from session
	//finding unpublished book id from 	database
	BL.Blist = model.GetBookList(1, Currentloggedin) // 1 - publishedbook, 0 - No specific user
	t, _ := template.ParseFiles("HTMLS/my-published-book.html")
	t.Execute(res, BL)
}

func MyUnpublishedBook(res http.ResponseWriter, req *http.Request) {
	var BL model.BookList
	log.Println("Method:MyUnpublishedBook -> Publisher id = ", Currentloggedin)
	//Take publisherid(Currentloggedin) from session
	//finding unpublished book id from 	database
	BL.Blist = model.GetBookList(0, Currentloggedin) // 1 - publishedbook, 0 - No specific user
	t, _ := template.ParseFiles("HTMLS/my-unpublished-book.html")
	t.Execute(res, BL)
}

//publishing a new Book
func PublishNewBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Method:PublisNewBook", req.Method)

	if req.Method == "POST" {

		//finding unique book id
		var bid int
		bid = model.GenerateID(2)
		var book_id string
		book_id = strconv.Itoa(bid)

		//finding book publisher id
		publisher_id := Currentloggedin //it is temporary finally session will generat publisher_id

		//finding book title,description and isbn no
		title := req.FormValue("title")
		description := req.FormValue("description")
		isbn := req.FormValue("isbn")

		//finding book cover_photo and pdf version of book
		file, handler, err := req.FormFile("cover_photo")
		file2, handler2, err2 := req.FormFile("pdf")

		//error checking
		if err != nil {
			fmt.Println(err)
			http.Redirect(res, req, "/publish-new-book", 301)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			http.Redirect(res, req, "/publish-new-book", 301)
			return
		}

		//closing
		defer file.Close()
		defer file2.Close()

		//changing file name
		handler.Filename = book_id + ".jpg"
		handler2.Filename = book_id + ".pdf"
		log.Println("File Name ", handler.Filename)
		log.Println("Pdf Name ", handler2.Filename)

		//saving file to their destination
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"CoverPhoto"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		f2, err2 := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"Pdf"+string(os.PathSeparator)+handler2.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			http.Redirect(res, req, "/publish-new-book", 301)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			http.Redirect(res, req, "/publish-new-book", 301)
			return
		}
		defer f.Close()
		defer f2.Close()
		io.Copy(f, file)
		io.Copy(f2, file2)

		fmt.Println("Book id ", bid, "Publisher ", publisher_id, " title ", title, "Description ", description, "ISBN ", isbn)
		//	db.Exec("INSERT INTO Book (Book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published,Average_rating) VALUES (?,?,?,?,?,?,?,?,?)", cnt, publisher_id, title, description, handler.Filename, isbn, 0, 0.0)

		//if isbn number is not uniqu then
		//db.QueryRow("SELECT Isbn  FROM Book WHERE Isbn=?", isbn).Scan(&isbnexist)
		tmpBook := model.GetBookByIsbn(isbn)
		if tmpBook.Isbn == isbn {
			fmt.Println("Isbn already used")
			http.Redirect(res, req, "/publish-new-book", 301)
			return
		}

		//value updated to database
		//db.Exec("INSERT INTO  Book (book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published, Average_rating) VALUES (?, ?, ?,? , ?, ?, ?, 0, 0)", cnt, publisher_id, title, description, handler.Filename, isbn, handler2.Filename) //, cnt, publisher_id, title, description, handler.Filename, isbn, handler2.Filename, 0, 0)
		var book model.Book
		book.Set(bid, publisher_id, title, description, handler.Filename, isbn, handler2.Filename, 0, 0)
		model.SetBook(book)
		fmt.Println("New Book Store successfully")
		http.Redirect(res, req, "/publish-new-book", 301)
	} else {
		t, _ := template.ParseFiles("HTMLS/publish-new-book.html")
		t.Execute(res, nil)
	}
}

//publisher update info of his book waiting for admin approval
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var book_id = req.URL.Query().Get("book")
	bid, _ := strconv.Atoi(book_id)
	var TmpBook model.BookP
	TmpBook = model.GetBook(bid)
	t, _ := template.ParseFiles("HTMLS/update-book.html")

	if req.Method != http.MethodPost {
		fmt.Println("Method:UpdateBook GET Method, redirect from : /my-unpublished-book")
		t.Execute(res, TmpBook)
		return
	}

	fmt.Println("Method: UpdateBook  POST Method,  redirect from : /update-book")
	//starting upload cover
	file, handler, err := req.FormFile("cover_photo")
	//error checking
	if err != nil {
		fmt.Println("No cover photo")
		fmt.Println(err)
	} else {
		fmt.Println("New cover photo found :", handler.Filename)
		//closing
		defer file.Close()
		//changing file name
		handler.Filename = book_id + ".jpg"
		fmt.Println("Cover Name ", handler.Filename)
		//saving file to their destination and at first deleting the already existing file
		os.Remove("." + string(os.PathSeparator) + "uploads" + string(os.PathSeparator) + "CoverPhoto" + string(os.PathSeparator) + handler.Filename)
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"CoverPhoto"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} //end of uploading cover photo

	//update pdf file
	file, handler, err = req.FormFile("pdf")
	//error checking
	if err != nil {
		fmt.Println("No cover photo")
		fmt.Println(err)

	} else {
		fmt.Println("New pdf found :", handler.Filename)
		defer file.Close()
		//changing file name
		handler.Filename = book_id + ".pdf"
		fmt.Println("Pdf Name ", handler.Filename)
		//saving file to their destination and at first deleting the already existing file
		os.Remove("." + string(os.PathSeparator) + "uploads" + string(os.PathSeparator) + "Pdf" + string(os.PathSeparator) + handler.Filename)
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"Pdf"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} //end of uploading pdf file

	title := req.FormValue("title")
	if title != "" {
		fmt.Println(" title update : ", title)
		model.UpdateBookTitle(bid, title)
	}

	description := req.FormValue("description")
	if description != "" {
		fmt.Println(" description update : ", description)
		model.UpdateBookDescription(bid, description)
	}

	t.Execute(res, TmpBook)
}

func ViewBook(res http.ResponseWriter, req *http.Request) {

	var book_id = req.URL.Query().Get("book")
	fmt.Println("Requested book ID : ", book_id)
	bid, _ := strconv.Atoi(book_id)
	var b model.BookP

	//db.QueryRow("select pdf,Title,cover_photo,description,Isbn,Average_rating name from Book,user_info where  user_info.user_id=Book.publisher_id and is_published=1 and Book.book_id=?", book_id).Scan(&pdf, &bname, &cover_photo, &description, &isbn, &average_rating, &pname)
	b = model.GetBook(bid)
	fmt.Println("Single book view ViewBook.go")
	t, _ := template.ParseFiles("HTMLS/view-book.html")
	t.Execute(res, b)
}
