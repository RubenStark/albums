package database

import (
	//importamos el paquete GORM
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	//Se define una variable global DBConn que se 
	//utiliza para conectarse a la base de datos SQLite.
	DBConn *gorm.DB
)
