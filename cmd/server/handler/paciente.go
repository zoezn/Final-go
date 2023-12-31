package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/internal/paciente"
	"github.com/zoezn/Final-go/pkg/web"
)

type pacienteHandler struct {
	s paciente.PacienteService
}

func NewPacienteHandler(s paciente.PacienteService) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// @Summary Actualizar paciente por Id.
// @Tags Paciente
// @Produce json
// @Success 200 {object} web.Response
// @Router /paciente/:id [get]
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
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

// @Summary Crear nuevo paciente.
// @Tags Paciente
// @Produce json
// @Success 200 {object} web.Response
// @Router /paciente/ [post]
func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := validateEmptyPacientes(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// @Summary Eliminar paciente por Id.
// @Tags Paciente
// @Produce json
// @Success 200 {object} web.Response
// @Router /pacientes/:id [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {
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

// @Summary Actualizar paciente por Id.
// @Tags Paciente
// @Produce json
// @Success 200 {object} web.Response
// @Router /pacientes/:id [put]
func (h *pacienteHandler) Put() gin.HandlerFunc {
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
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := validateEmptyPacientes(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// @Summary Actualizar campos deseados de paciente por Id.
// @Tags Paciente
// @Produce json
// @Success 200 {object} web.Response
// @Router /pacientes/:id [patch]
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		// Id        int    `json:"id"`
		Nombre    string `json:"nombre" binding:"-"`
		Apellido  string `json:"apellido" binding:"-"`
		Domicilio string `json:"domicilio" binding:"-"`
		DNI       int    `json:"DNI" binding:"-"`
		Alta      string `json:"alta" binding:"-"`
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
		update := domain.Paciente{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Domicilio: r.Domicilio,
			DNI:       r.DNI,
			Alta:      r.Alta,
		}

		paciente, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, paciente)
	}
}
