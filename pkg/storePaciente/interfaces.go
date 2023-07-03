package storePaciente

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type PacienteInterface interface {
	Read(id int) (domain.Paciente, error)
	ReadByDNI(dni int) (domain.Paciente, error)
	Create(paciente domain.Paciente) error
	Update(id int, paciente domain.Paciente) error
	Delete(id int) error
	Exists(id int) bool
	ExistsByDni(dni int) bool
}
