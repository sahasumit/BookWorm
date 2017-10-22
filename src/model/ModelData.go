package model

type User struct {
	UserId, IsActive                int
	Email, Password, Name, UserType string
}

func (u *User) Set(uid int, email string, pass string, name string, isActv int, userType string) {
	u.UserId = uid
	u.Email = email
	u.Name = name
	u.Password = pass
	u.IsActive = isActv
	u.UserType = userType
}

type Book struct {
	BookId, PubId, IsPubed               int
	Title, Description, Cover, Isbn, Pdf string
	AvrgRating                           float32
}

func (b *Book) Set(bookId int, pubId int, title_ string, descrptn string, cover string, isbn string, pdf string, isPubed int, avrgRat float32) {
	b.BookId = bookId
	b.PubId = pubId
	b.Title = title_
	b.Description = descrptn
	b.Cover = cover
	b.Isbn = isbn
	b.Pdf = pdf
	b.IsPubed = isPubed
	b.AvrgRating = avrgRat
}

type BookP struct {
	BookId, PubId, IsPubed               int
	Title, Description, Cover, Isbn, Pdf string
	AvrgRating                           float32
	PubName                              string
}

func (b *BookP) Set(bookId int, pubId int, title_ string, descrptn string, cover string, isbn string, pdf string, isPubed int, avrgRat float32, pubName string) {
	b.BookId = bookId
	b.PubId = pubId
	b.Title = title_
	b.Description = descrptn
	b.Cover = cover
	b.Isbn = isbn
	b.Pdf = pdf
	b.IsPubed = isPubed
	b.AvrgRating = avrgRat
	b.PubName = pubName
}

type BookList struct {
	Blist []BookP
	//Pub   string //publisher name
}
