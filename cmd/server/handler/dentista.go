package handler

import (
	"errors"
	"os"
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

func TokenAuthMiddleware(token string) gin.HandlerFunc {
	requiredToken := os.Getenv("Token ")
	return func(c *gin.Context) {
		// token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != requiredToken {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		c.Next()
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

func validateEmptys(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Nombre == "" || dentista.Apellido == "" || dentista.Matricula == "":
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}

// Post crea un nuevo dentista
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid json"))
			return
		}
		valid, err := validateEmptys(&dentista)
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

func (h *dentistaHandler) Put() gin.HandlerFunc {
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
		valid, err := validateEmptys(&dentista)
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
