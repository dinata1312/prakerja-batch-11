package controllers

import (
	"fmt"
	"net/http"
	"prakerja/model"

	"github.com/gin-gonic/gin"
)

// books adalah slice untuk menyimpan data buku (instances model.Book).
var books = []model.Book{}

// CreateBook menangani pembuatan buku baru.
func CreateBook(ctx *gin.Context) {
	// Deklarasi variabel newBook dengan tipe model.Book.
	var newBook model.Book

	// Coba ikat data JSON yang masuk ke variabel newBook.
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		// Jika ada kesalahan dalam ikatan JSON, kirim respons Bad Request.
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Menghasilkan BookID unik untuk buku baru dan menambahkannya ke slice books.
	newBook.BookID = fmt.Sprintf("%d", len(books)+1)
	books = append(books, newBook)

	// Kirim respons dengan buku yang telah dibuat sebagai JSON.
	ctx.JSON(http.StatusCreated, gin.H{"book": newBook})
}

// UpdateBook menangani pembaruan buku yang sudah ada berdasarkan BookID.
func UpdateBook(ctx *gin.Context) {
	// Ambil BookID dari parameter URL.
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook model.Book

	// Coba ikat data JSON yang masuk ke variabel updatedBook.
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		// Jika ada kesalahan dalam ikatan JSON, kirim respons Bad Request.
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Iterasi melalui slice books untuk menemukan dan memperbarui buku yang cocok.
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			books[i] = updatedBook
			books[i].BookID = bookID
			break
		}
	}

	// Jika tidak ditemukan buku dengan ID yang cocok, kirim respons Not Found.
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
		})
		return
	}

	// Kirim respons berhasil sebagai JSON.
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s has been successfully updated", bookID),
	})
}

// GetBook mengambil informasi buku berdasarkan BookID.
func GetBook(ctx *gin.Context) {
	// Ambil BookID dari parameter URL.
	bookID := ctx.Param("bookID")
	condition := false
	var bookData model.Book

	// Iterasi melalui slice books untuk mencari buku yang cocok.
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			bookData = books[i]
			break
		}
	}

	// Jika tidak ditemukan buku dengan ID yang cocok, kirim respons Not Found.
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
		})
		return
	}

	// Kirim informasi buku yang ditemukan sebagai JSON.
	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

// DeleteBook menghapus buku berdasarkan BookID.
func DeleteBook(ctx *gin.Context) {
	// Ambil BookID dari parameter URL.
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	// Iterasi melalui slice books untuk mencari buku yang cocok.
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	// Jika tidak ditemukan buku dengan ID yang cocok, kirim respons Not Found.
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
		})
		return
	}

	// Hapus buku dari slice books.
	copy(books[bookIndex:], books[bookIndex+1:])
	books[len(books)-1] = model.Book{}
	books = books[:len(books)-1]

	// Kirim respons berhasil sebagai JSON.
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s has been successfully deleted", bookID),
	})
}

// GetAllBooks mengambil semua buku dalam koleksi.
func GetAllBooks(ctx *gin.Context) {
	// Kirim semua buku dalam koleksi sebagai JSON.
	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
