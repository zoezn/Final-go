package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/internal/turno"
	"github.com/zoezn/Final-go/pkg/web"
)

type turnoHandler struct {
	s turno.TurnoService
}

func NewTurnoHandler(s turno.TurnoService) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// supuestamente esto es para documentar swagger
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

func validateEmptyTurnos(turno *domain.Turno) (bool, error) {
	switch {
	case turno.Dentista.Nombre == "" || turno.Paciente.Nombre == "" || turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "":
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}

func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := validateEmptyTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := validateEmptyTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Id          int             `json:"id"`
		Paciente    domain.Paciente `json:"paciente" binding:"required"`
		Dentista    domain.Dentista `json:"dentista" binding:"required"`
		Fecha       string          `json:"fecha" binding:"required"`
		Hora        string          `json:"hora" binding:"required"`
		Descripcion string          `json:"descripcion" binding:"optional"`
	}

	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		update := domain.Turno{
			Dentista:    r.Dentista,
			Paciente:    r.Paciente,
			Fecha:       r.Fecha,
			Hora:        r.Hora,
			Descripcion: r.Descripcion,
		}

		turno, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, turno)
	}
}
