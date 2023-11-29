package main

import (
	"database/sql"
	"net/http"
	"time"
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

type Person struct
{
    Id        string       `json:"id"`
    Name      string       `json:"name"`
    Email     string       `json:"email"`
    CreatedAt time.Time    `json:"createdAt"`
    UpdatedAt sql.NullTime `json:"updatedAt"`
}

var personList = []Person 	{
    {
        Id: "3ba86131-36cf-4052-b312-11ddbaa8fe9f",
        Name: "One",
        Email: "one@email.com",
        CreatedAt: time.Now(),
        UpdatedAt: sql.NullTime {
            Time: time.Now(),
            Valid: true,
        },
    },
    {
        Id: "6dca7cd3-87a9-497d-8cbc-4250872c65a0",
        Name: "Two",
        Email: "two@email.com",
        CreatedAt: time.Now(),
        UpdatedAt: sql.NullTime {},
    },
}

func getPeople(c *gin.Context) {
    c.JSON(http.StatusOK, personList)
}

func addPerson(c *gin.Context) {
    var requestPerson Person

    if err := c.BindJSON(&requestPerson); err != nil {
        fmt.Print(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    personList = append(personList, requestPerson)
    c.JSON(http.StatusCreated, nil)
}

func createRoutes(router *gin.Engine) {
	router.GET("/people", getPeople)
	router.POST("/people", addPerson)
}

func main() {
	router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{ "http://localhost:8080", "http://127.0.0.1:8080" }
    config.AllowMethods = []string{ "GET", "POST", "PUT", "DELETE" }
    router.Use(cors.New(config))

	createRoutes(router)

    err := router.Run("localhost:8082")
    
    if err != nil {
        fmt.Printf("An error occurred: %+v\n", err)
    }
}
