package handler

import (
	"errors"
	"fmt"
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

// @Summary Buscar turno por Id.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turnos/:id [get]
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		paciente, err := h.s.GetByID(id)
		fmt.Println(err)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// @Summary Obtener turno por dni de paciente.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turnos?dni [get]
func (h *turnoHandler) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		// dniParam := c.Param("dni")
		dniParam := c.Query("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid dni"))
			return
		}
		turno, err := h.s.GetByDNI(dni)
		fmt.Println(err)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno not found"))
			return
		}
		web.Success(c, 200, turno)
	}
}

// @Summary Crear nuevo turno.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turno/ [post]
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
		fmt.Println(err)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := ValidateEmptyTurnos(&turno)
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

// @Summary Crear turno por matricula dentista y id paciente.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turnos/noids [post]
func (h *turnoHandler) PostMatriculaId() gin.HandlerFunc {
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
		fmt.Println(err)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := ValidateEmptyTurnosNoIds(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.CreateWithoutIds(turno, turno.Matricula, turno.DNI)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// @Summary Eliminar turno por Id.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turnos/:id [delete]
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

// @Summary Actualizar turno por Id.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turnos/:id [put]
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
		valid, err := ValidateEmptyTurnos(&turno)
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

// @Summary Actualizar campos deseados de turno por Id.
// @Tags Turno
// @Produce json
// @Success 200 {object} web.Response
// @Router /turno/:id [patch]
func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Id          int             `json:"id"`
		Paciente    domain.Paciente `json:"paciente" binding:"-"`
		Dentista    domain.Dentista `json:"dentista" binding:"-"`
		Fecha       string          `json:"fecha" binding:"-"`
		Hora        string          `json:"hora" binding:"-"`
		Descripcion string          `json:"descripcion" binding:"-"`
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
