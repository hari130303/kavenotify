package main

import (
	"fmt"
	"kavenotify/controller"
	dbfunction "kavenotify/db_function"
	envconfig "kavenotify/env_config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment configuration
	err := envconfig.ENV_CONGIF()
	if err != nil {
		log.Fatalf("Error loading ENV config: %v", err)
	}
	fmt.Println("ENV config loaded")

	// Initialize database connection
	err = dbfunction.DB_conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connected")

	dbfunction.SchemaSetFunc()

	// Initialize mongo database connection
	err = dbfunction.MONOGO_DB_Conn()
	if err != nil {
		log.Fatalf("Error connecting to mongo database: %v", err)
	}
	fmt.Println("Mongo database connected")

	// Initialize rabbitmq
	err = dbfunction.RabbitQueue()
	if err != nil {
		log.Fatalf("Error creating rabbit queue: %v", err)
	}
	fmt.Println("Rabbit queue created")

	//start the rabbit mq listener 
	go dbfunction.RabbitReciever()

	// Initialize Gin router
	r := gin.Default()
	r.Use(corsMiddleware())

	// Template routes
	r.Any("/", controller.LoginTemplate)
	r.Any("/register", controller.RegisterTemplate)
	r.Any("/dashboard/single_message", controller.SingleMessage)
	r.Any("/dashboard/bulk_message", controller.BulkMessage)
	r.Any("/dashboard/customer_setup", controller.CustomerSetup)
	r.Any("/dashboard/number_manage", controller.NumberManage)
	r.Any("/dashboard/user_log", controller.UserLoginLog)

	// API call control
	r.POST("/login", controller.Login)
	r.POST("/register-api", controller.RegisterApi)
	r.POST("/get/customer_type", controller.GetCustomerType)
	r.POST("/send/single_message", controller.SendSingleMessage)
	r.POST("/get/customer_type/table", controller.GetCustomerTypeTable)
	r.POST("/get/number_manage/table", controller.NumberManageTable)
	r.POST("/get/user_login_log", controller.GetUserLoginLog)

	// Channel to listen for OS shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go dbfunction.SessionManager()
	// Start the Gin server in a separate goroutine
	go func() {
		fmt.Println("Starting server on", envconfig.HOST_PORT)
		if err := r.Run(envconfig.HOST_PORT); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	fmt.Println("\nShutting down server...")

	// Close database connection gracefully
	dbfunction.DBClose()

	fmt.Println("Server exited gracefully.")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http,https")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
