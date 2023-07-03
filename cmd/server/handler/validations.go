package handler

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zoezn/Final-go/internal/domain"
)

func ValidateToken(c *gin.Context) error {
	token := c.GetHeader("TOKEN")
	if token == "" {
		return errors.New("Token not found")
	}
	if token != os.Getenv("TOKEN") {
		return errors.New("Invalid token")
	}
	return nil
}

func ValidateEmptyDentista(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Nombre == "" || dentista.Apellido == "" || dentista.Matricula == "":
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}

func validateEmptyPacientes(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.Nombre == "" || paciente.Apellido == "" || paciente.Domicilio == "" || paciente.DNI < 0 || paciente.Alta == "":
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}

func ValidateEmptyTurnos(turno *domain.Turno) (bool, error) {
	switch {
	case turno.Dentista.Id == 0 || turno.Paciente.Id == 0 || turno.Fecha == "" || turno.Hora == "" || turno.Descripcion == "":
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}
