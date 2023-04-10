package database

import (
	"errors"
	"fmt"
	"log"
	"project-1-chapter-2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	port     = 5432
	dbname   = "db-gorm-book-sql"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}

func GetBookAllDB(book []models.Book) (bookDatas []models.Book, err error) {
	db := GetDB()

	err = db.Find(&book).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book data not found")
		err = errors.New("gagal mendapatkan data")
		return
	}

	bookDatas = append(bookDatas, book...)

	return
}

func GetBookIdDB(id int, book models.Book) (bookDatas models.Book, err error) {
	db := GetDB()

	err = db.First(&book, "id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	bookDatas = book
	return
}

func CreateBookDB(book models.Book) (bookDatas models.Book, err error) {
	var getID int
	var getBook models.Book
	db := GetDB()

	_ = db.Select("id").Last(&getBook).Scan(&getID)

	book = models.Book{
		ID:       uint(getID) + 1,
		NameBook: book.NameBook,
		Author:   book.Author,
	}

	err = db.Create(&book).Error
	if err != nil {
		return
	}

	bookDatas = book
	return
}

func UpdateBookDB(id int, book models.Book) (bookDatas models.Book, err error) {
	db := GetDB()

	var findBook models.Book

	err = db.Where("id = ?", id).First(&findBook).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book not found")
		err = errors.New("buku tidak ditemukan")
		return
	}

	err = db.Model(&book).Where("id = ?", id).Updates(models.Book{
		ID:        uint(id),
		NameBook:  book.NameBook,
		Author:    book.Author,
		UpdatedAt: findBook.UpdatedAt,
	}).Error

	if err != nil {
		log.Println("Error update book data", err)
		err = errors.New("error update book data")
		return
	}

	bookDatas = book

	return
}

func DeleteBookDB(id int) (err error) {
	db := GetDB()
	var book models.Book

	err = db.Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return
}
