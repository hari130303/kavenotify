package envconfig

import (
	"fmt"
	"os"
)

var DB_USER, DB_PASS, DB_HOST, DB_NAME, DB_PORT, HOST_PORT, AES32KEY, AESIVKEY, SCHEMA_NAME, MONGO_USER, MONGO_PASS, MONGO_PORT, MONGO_HOST, MONGO_DB, MONGO_USER_COLL, MONGO_LOGGER_COLL string

func ENV_CONGIF() error {
	// err := godotenv.Load()
	// if err != nil {
	// 	return err
	// }
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASSWD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	HOST_PORT = os.Getenv("HOST_PORT")
	AES32KEY = os.Getenv("AES32KEY")
	AESIVKEY = os.Getenv("AESIVKEY")
	SCHEMA_NAME = os.Getenv("SCHEMA_NAME")
	MONGO_HOST = os.Getenv("MONGO_HOST")
	MONGO_PORT = os.Getenv("MONGO_PORT")
	MONGO_USER = os.Getenv("MONGO_USER")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	MONGO_DB = os.Getenv("MONGO_DB")
	MONGO_USER_COLL = os.Getenv("MONGO_USER_COLL")
	MONGO_LOGGER_COLL = os.Getenv("MONGO_LOGGER_COLL")
	fmt.Println("config schema :", SCHEMA_NAME)

	return nil
}
