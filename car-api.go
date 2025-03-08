package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type car struct {
    ID     string  `json:"id"`
    Name   string  `json:"Name"`
}

var cars = []car{
    {ID: "1", Name: "special car"},
    {ID: "2", Name: "another special car"},
}

func main() {
    router := gin.Default()
    router.GET("/cars", getCars)
    router.GET("/cars/:id", getCarById)
    router.POST("/cars", postCar)

    router.Run("localhost:8080")
}

func getCars(c *gin.Context){
    c.IndentedJSON(http.StatusOK, cars)
}

func getCarById(c *gin.Context){
    id := c.Param("id")

    for _, a := range cars {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found! :("})
}

func postCar(c *gin.Context){
    var newCar car

    if err := c.BindJSON(&newCar); err != nil {
        return
    }

    cars = append(cars, newCar)
    c.IndentedJSON(http.StatusCreated, newCar)
}