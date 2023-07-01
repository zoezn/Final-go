package storePaciente

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type PacienteInterface interface {
	Read(id int) (domain.Paciente, error)
	Create(product domain.Paciente) error
	Update(product domain.Paciente) error
	Delete(id int) error
	Exists(id int) bool
}
