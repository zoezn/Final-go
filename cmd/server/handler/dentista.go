package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoezn/Final-go/internal/dentista"
	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/pkg/web"
)

type dentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Dentista not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

// Post crea un nuevo dentista
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ValidateToken(c); err != nil {
			web.Failure(c, 400, err)
			return
		}

		var dentista domain.Dentista

		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := ValidateEmptyDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

func (h *dentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ValidateToken(c); err != nil {
			web.Failure(c, 400, err)
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

func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ValidateToken(c); err != nil {
			web.Failure(c, 400, err)
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
			web.Failure(c, 404, errors.New("Dentista not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := ValidateEmptyDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido" binding:"-"`
		Nombre    string `json:"nombre" binding:"-"`
		Matricula string `json:"matricula" binding:"-"`
	}
	return func(c *gin.Context) {
		if err := ValidateToken(c); err != nil {
			web.Failure(c, 400, err)
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
			web.Failure(c, 404, errors.New("Dentista not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}

		update := domain.Dentista{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}

		dentista, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, dentista)
	}
}
