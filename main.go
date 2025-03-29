package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"crud/database"
	"crud/http"
	"crud/logic"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
    logger := logrus.New()
    logger.SetFormatter(&logrus.TextFormatter{})
    logger.SetLevel(logrus.InfoLevel)

    conn, err := dbr.Open("postgres", "user=postgres password=postgres dbname=person sslmode=disable", nil)
    
    if err != nil {
        log.Fatalf("DB connection error: %v", err)
    }

    session := conn.NewSession(nil)
    repo := database.NewPersonRepository(session)
    personLogic := logic.NewPersonLogic(repo)
    personHandler := http.NewPersonHandler(personLogic, logger)
    e := echo.New()

    e.GET("/person", personHandler.GetAllPersons)
    e.GET("/person/:id", personHandler.GetPersonByID)
    e.POST("/person", personHandler.CreatePerson)
    e.PUT("/person/:id", personHandler.UpdatePerson)
    e.DELETE("/person/:id", personHandler.DeletePerson)

    port := 8080
    logger.Infof("Server has been started port %d", port)
    
    if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
        logger.Fatalf("Server start error: %v", err)
    }
}
