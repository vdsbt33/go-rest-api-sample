package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func getPeople(c *gin.Context) {
    result := []Person 	{
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

    c.JSON(http.StatusOK, result)
}

func createRoutes(router *gin.Engine) {
	router.GET("/people", getPeople)
}

func main() {
	router := gin.Default()
	createRoutes(router)

	router.Run("localhost:8080")
}
