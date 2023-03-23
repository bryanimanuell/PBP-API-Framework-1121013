package main

import (
	"src/GIN/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.InsertUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)

	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	// router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	router.Run("localhost:8080")ab
}
