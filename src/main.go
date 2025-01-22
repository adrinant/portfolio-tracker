package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"portfolio-tracker/controller"
	"portfolio-tracker/repository"
	"portfolio-tracker/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow all origins for CORS
		c.Header("Access-Control-Allow-Origin", "*")
		// Allow credentials
		c.Header("Access-Control-Allow-Credentials", "true")
		// Allow specific headers
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// Allow specific methods, including DELETE
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
		// Allow the server to vary based on the origin
		c.Header("Vary", "Origin")

		// Handle preflight requests (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		// Continue processing the request
		c.Next()
	}
}

func main() {
	log.Println("Starting Server ...")

	db, err := initDB()

	if err != nil {
		log.Fatalf("Unable to initialize database: %v\n", err)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepository)

	controller.NewController(&controller.Config{
		R:                  router,
		TransactionService: transactionService,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		//service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful server shutdown
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := db.close(); err != nil {
		log.Fatalf("A problem occured gracefully shutting down the database connection: %v\n", err)
	}

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
