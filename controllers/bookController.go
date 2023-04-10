package controllers

import (
	"fmt"
	"log"
	"net/http"
	"project-1-chapter-2/database"
	"project-1-chapter-2/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookAll(ctx *gin.Context) {
	var bookDatas []models.Book

	getBooks, err := database.GetBookAllDB(bookDatas)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Gagal mendapatkan request %+v", err.Error()),
		})
		return
	}

	if getBooks == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": getBooks,
	})
}

func GetBookId(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var bookDatas models.Book

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal mengconvert id")
		return
	}

	book, err := database.GetBookIdDB(convID, bookDatas)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convID),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	var bookDatas models.Book

	if err := ctx.ShouldBindJSON(&bookDatas); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := database.CreateBookDB(bookDatas)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Gagal membuat buku %+v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookDatas models.Book

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal mengconvert id")
		return
	}

	if err := ctx.ShouldBindJSON(&bookDatas); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := database.UpdateBookDB(convID, bookDatas)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Gagal update buku %+v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal mengconvert id")
		return
	}

	err = database.DeleteBookDB(convID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d successfully deleted", convID),
	})

}
