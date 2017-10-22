package model

import (
	"log"
	"model/dbcon"
	"strconv"
)

func SetUser(user User) {
	log.Println("User ", user.Name)
	_, err := dbcon.Db.Exec("INSERT INTO user_info (user_id, email, password, name, is_active, user_type) VALUES (?, ?, ?, ?, ?, ?)", user.UserId, user.Email, user.Password, user.Name, user.IsActive, user.UserType)
	if err != nil {
		log.Print(err)
	}

}
func GetUser(Email string) User {
	var user User
	err := dbcon.Db.QueryRow("SELECT * FROM user_info WHERE email=?", Email).Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.IsActive, &user.UserType)
	log.Println(" UserInfo : ", user.Email, " ", user.Name)
	if err != nil {
		log.Println(err)
	}
	return user
}
func GetUserList() []User {
	var u User
	var ul []User
	rows, err := dbcon.Db.Query("SELECT * FROM user_info WHERE user_type != 'admin'")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.Email, &u.Password, &u.Name, &u.IsActive, &u.UserType)
		if err != nil {
			log.Println(err)
		} else {
			ul = append(ul, u)
		}
	}
	return ul
}

func GenerateID(TableType int) int {
	var tablename string
	if TableType == 1 {
		tablename = "user_info"
	}
	if TableType == 2 {
		tablename = "Book"
	}
	var xid int
	var sql = "select count(*) from " + tablename
	dbcon.Db.QueryRow(sql).Scan(&xid)
	log.Print("Generate user id: ", xid, tablename)
	xid += 1
	return xid
}

/*Get a List of book in a Array of BookP
bookType - (1/0) ->> (published/unpublished)
pubId - 0 ->> for no specific publisher_id but any publisher_id
pubId - greater than 0 and also matching publisher_id ->> for specific publishers book
*/

func GetBookList(bookType int, pubId int) []BookP {

	sql := "select * from Book where is_published = " + strconv.Itoa(bookType)
	if pubId != 0 {
		sql += " AND publisher_id = " + strconv.Itoa(pubId)
	}
	var bookArray []BookP
	rows, err := dbcon.Db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var TmpBook BookP
		err := rows.Scan(&TmpBook.BookId, &TmpBook.PubId, &TmpBook.Title, &TmpBook.Description, &TmpBook.Cover, &TmpBook.Isbn, &TmpBook.Pdf, &TmpBook.IsPubed, &TmpBook.AvrgRating)
		if err != nil {
			log.Println(err)
		}
		//to get Publisher name from user_info Table using PubId from Book Table
		sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(TmpBook.PubId)
		dbcon.Db.QueryRow(sql).Scan(&TmpBook.PubName)

		bookArray = append(bookArray, TmpBook)
		log.Println("Books ID ", TmpBook.BookId, " book name ", TmpBook.Title)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return bookArray
}

func GetBook(bookId int) BookP {
	var b BookP
	sql := "select * from Book where book_id = " + strconv.Itoa(bookId)
	dbcon.Db.QueryRow(sql).Scan(&b.BookId, &b.PubId, &b.Title, &b.Description, &b.Cover, &b.Isbn, &b.Pdf, &b.IsPubed, &b.AvrgRating)
	sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(b.PubId)
	dbcon.Db.QueryRow(sql).Scan(&b.PubName)
	return b
}

func GetBookByIsbn(isbn string) BookP {
	var b BookP
	sql := "select * from Book where Isbn = " + isbn
	dbcon.Db.QueryRow(sql).Scan(&b.BookId, &b.PubId, &b.Title, &b.Description, &b.Cover, &b.Isbn, &b.Pdf, &b.IsPubed, &b.AvrgRating)
	//sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(b.PubId)
	//dbcon.Db.QueryRow(sql).Scan(&b.PubName)
	return b
}

func SetBook(book Book) {
	_, err := dbcon.Db.Exec("INSERT INTO  Book (book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published, Average_rating) VALUES (?, ?, ?,? , ?, ?, ?, ?, ?)", book.BookId, book.PubId, book.Title, book.Description, book.Cover, book.Isbn, book.Pdf, book.IsPubed, book.AvrgRating)
	if err != nil {
		log.Println(err)
	}
}

func UpdateBookTitle(bookId int, bookTitle string) {
	dbcon.Db.Exec("UPDATE Book SET Title=? WHERE book_id=?", bookTitle, bookId)
}

func UpdateBookDescription(bookId int, bookDescrptn string) {
	dbcon.Db.Exec("UPDATE Book SET description=? WHERE book_id=?", bookDescrptn, bookId)
}

/*publish, unpublish, rejec a book with id of bookId
isPub = 0 >> unpublish
isPub = 1 >> publish
isPub = 2 >> reject
*/
func PublishBook(bookId int, isPub int) {
	dbcon.Db.Exec("UPDATE Book SET is_published=? WHERE book_id=?", isPub, bookId)
}
