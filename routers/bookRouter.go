package routers

import (
	"prakerja/controllers" // Mengimpor modul 'controllers' untuk mengakses fungsi controller.

	"github.com/gin-gonic/gin" // Mengimpor modul 'gin' untuk membuat router web.
)

// StartServer adalah fungsi yang memulai server web menggunakan Gin.
func StartServer() *gin.Engine {

	// Membuat instance router menggunakan default Gin.
	router := gin.Default()

	// Menentukan rute-rute HTTP dan mengaitkannya dengan fungsi-fungsi controller yang sesuai.

	router.POST("/books", controllers.CreateBook)           // Rute HTTP POST untuk membuat buku baru, menggunakan fungsi CreateBook dari controllers.
	router.PUT("/books/:bookID", controllers.UpdateBook)    // Rute HTTP PUT untuk memperbarui buku berdasarkan BookID, menggunakan fungsi UpdateBook dari controllers.
	router.GET("/books/:bookID", controllers.GetBook)       // Rute HTTP GET untuk mengambil informasi buku berdasarkan BookID, menggunakan fungsi GetBook dari controllers.
	router.GET("/books", controllers.GetAllBooks)           // Rute HTTP GET untuk mengambil semua buku dalam koleksi, menggunakan fungsi GetAllBooks dari controllers.
	router.DELETE("/books/:bookID", controllers.DeleteBook) // Rute HTTP DELETE untuk menghapus buku berdasarkan BookID, menggunakan fungsi DeleteBook dari controllers.

	// Mengembalikan instance router yang telah dikonfigurasi.
	return router
}
