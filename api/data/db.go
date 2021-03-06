package data

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/yeadan/okupa/api/models"
)

// InitEnv Inicializa el .env para poder leerlo (*NO - y creamos directorios
//para images por si no existen (si existen no hace nada) *)
func InitEnv() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	//os.Mkdir("./images",0777)
	//os.Mkdir("./images/avatars",0777)
}

// ConnectDB Crea la conexión a la base de datos postgres
func ConnectDB() (*gorm.DB, error) {
	InitEnv()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbDriver := os.Getenv("DB_DRIVER")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	return gorm.Open(dbDriver, dbURI)
}

// InitDB - Inicializa la base de datos
func InitDB() {
	dbCnx, err := ConnectDB()
	if err == nil {
		defer dbCnx.Close()
		dbCnx.BlockGlobalUpdate(true)
		dbCnx.AutoMigrate(models.User{}, models.Okupa{}, models.UserOkupa{}, models.Owner{}, models.Property{})
		dbCnx.BlockGlobalUpdate(true)
	} else {
		panic(err)
	}
}
