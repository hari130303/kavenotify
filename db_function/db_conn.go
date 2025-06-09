package dbfunction

import (
	"context"
	"database/sql"
	"fmt"
	envconfig "kavenotify/env_config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "honnef.co/go/tools/config"
)

var db *sql.DB

var umdb *mongo.Collection
var lmdb *mongo.Collection

// username:=os.Genenv("sdclskm")

// var username, password, host, port, database string
var SchemaName string

func DB_conn() error {

	var err error
	// build the DSN
	// dsn := fmt.Sprintf("root:12345@tcp(localhost:3306)/kavenotify_db")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", envconfig.DB_USER, envconfig.DB_PASS, envconfig.DB_HOST, envconfig.DB_PORT, envconfig.DB_NAME)
	fmt.Println(dsn)
	// Open the connection
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	// close connection(s) when the surrounding function exits
	return err
}

func DBClose() {
	db.Close()
}

func SchemaSetFunc() {
	SchemaName = envconfig.SCHEMA_NAME
	fmt.Println("schema set : ", SchemaName)
}

func MONOGO_DB_Conn() error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure resources are cleaned up

	// MongoDB URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", envconfig.MONGO_USER, envconfig.MONGO_PASS, envconfig.MONGO_HOST, envconfig.MONGO_PORT)

	// fmt.Println(" mongo url :", uri)
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// clientOptions := options.Client().ApplyURI("mongodb://admin:12345@mongodb:27017/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("MongoDB connection error: %w", err)
	}

	// Ping to test the connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("MongoDB ping failed: %w", err)
	}

	// Initialize the global collection
	umdb = client.Database(envconfig.MONGO_DB).Collection(envconfig.MONGO_USER_COLL)
	lmdb = client.Database(envconfig.MONGO_DB).Collection(envconfig.MONGO_LOGGER_COLL)

	// log.Println("✅ MongoDB connected successfully")
	return nil
}
