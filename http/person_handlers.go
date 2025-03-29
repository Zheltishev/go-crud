package http

import (
	"net/http"
	"strconv"

	"crud/logic"
	"crud/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PersonHandler struct {
	Logic  *logic.PersonLogic
	logger *logrus.Logger
}

func NewPersonHandler(logic *logic.PersonLogic, logger *logrus.Logger) *PersonHandler {
	return &PersonHandler{
		Logic:  logic,
		logger: logger,
	}
}

func (h *PersonHandler) GetAllPersons(c echo.Context) error {
	h.logger.Info("Handling GetAllPersons request")
	
	persons, err := h.Logic.GetAllPersons()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get all persons")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	h.logger.WithField("count", len(persons)).Info("Successfully retrieved all persons")
	return c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPersonByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.WithField("id", id).Info("Handling GetPersonByID request")
	
	person, err := h.Logic.GetPersonByID(id)
	if err != nil {
		h.logger.WithError(err).WithField("id", id).Error("Person not found")
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Person not found"})
	}
	
	h.logger.WithField("id", id).Info("Successfully retrieved person")
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	h.logger.Info("Handling CreatePerson request")
	
	var person model.Person
	if err := c.Bind(&person); err != nil {
		h.logger.WithError(err).Error("Invalid input for person creation")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.Logic.CreatePerson(&person); err != nil {
		h.logger.WithError(err).Error("Failed to create person")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	h.logger.WithField("person_id", person.ID).Info("Successfully created person")
	return c.JSON(http.StatusCreated, person)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.WithField("id", id).Info("Handling UpdatePerson request")
	
	var person model.Person
	if err := c.Bind(&person); err != nil {
		h.logger.WithError(err).WithField("id", id).Error("Invalid input for person update")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.Logic.UpdatePerson(id, &person); err != nil {
		h.logger.WithError(err).WithField("id", id).Error("Failed to update person")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	h.logger.WithField("id", id).Info("Successfully updated person")
	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.WithField("id", id).Info("Handling DeletePerson request")
	
	if err := h.Logic.DeletePerson(id); err != nil {
		h.logger.WithError(err).WithField("id", id).Error("Failed to delete person")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	h.logger.WithField("id", id).Info("Successfully deleted person")
	return c.NoContent(http.StatusNoContent)
}
