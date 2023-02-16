package album

import (
	"fmt"

	// database variable es exportada por el paquete database
	"github.com/RubenStark/albums/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Recuerda que el nombre de la funcion debe
// Empezar con mayuscula para que sea publica

// GetAlbums returns all albums in the database
func GetAlbums(c *fiber.Ctx) {
	// Conecta con la base de datos
	db := database.DBConn
	// Inicia un array vacio de albums
	var albums []Album
	// Encuentra todos los albums y los guarda en el array
	db.Find(&albums)
	// Envia el array como JSON
	c.JSON(albums)

	fmt.Println("Todos los albums")
}

func GetAlbum(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var album Album
	db.Find(&album, id)
	c.JSON(album)

	fmt.Println("Un album en especifico")
}

func AddAlbum(c *fiber.Ctx) {
	db := database.DBConn

	album := new(Album)
	if err := c.BodyParser(album); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&album)
	c.JSON(album)

	fmt.Println("Añadir un album")
}

func DeleteAlbum(c *fiber.Ctx) { // Recibe el contexto de la petición
	id := c.Params("id")  // Obtiene el id del album a eliminar
	db := database.DBConn // Obtiene la conexión a la base de datos

	var album Album // Crea una instancia de Album

	// Busca el album con el id especificado. Si no lo encuentra, 
	//responde con un error.
	if result := db.First(&album, id); result.Error != nil {
		c.Status(500).Send("No album found with given ID")
		return
	}
	// Si el album existe lo elimina
	db.Delete(&album)
	// Responde con un mensaje indicando que el album fue eliminado
	c.Send("Album successfully deleted")

	fmt.Println("Eliminar un album")
}
