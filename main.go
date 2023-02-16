package main

import (
	"fmt"

	"github.com/RubenStark/albums/album"
	"github.com/RubenStark/albums/database"

	//go get -u github.com/gofiber/fiber
	"github.com/gofiber/fiber"

	//go get -u github.com/jinzhu/gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", album.GetAlbums)
	app.Get("/:id", album.GetAlbum)
	app.Post("/", album.AddAlbum)
	app.Delete("/:id", album.DeleteAlbum)
}

// Conenctar con la database
func initDatabase() {

	var err error

	//database.DBConn es exportada por el paquete database
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	fmt.Println("Database Conectada")

	// Migra la database
	database.DBConn.AutoMigrate(&album.Album{})

	fmt.Println("Database Migrada")

}

func main() {
	// Iniciar una nueva Fiber instance
	app := fiber.New()

	// Iniciar la conexion con la db
	initDatabase()
	// Cerrar la conexion con la db al terminar el programa
	defer database.DBConn.Close()

	// Configurar las rutas
	setupRoutes(app)

	// Empezar el server on port 3000
	app.Listen(3000)
}
