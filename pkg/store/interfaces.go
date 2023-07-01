package store

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type StoreInterface interface {
	Read(id int) (domain.Dentista, error)
	Create(product domain.Dentista) error
	Update(product domain.Dentista) error
	Delete(id int) error
	Exists(codeValue string) bool
}

type PacienteInterface interface {
	Read(id int) (domain.Paciente, error)
	Create(product domain.Paciente) error
	Update(product domain.Paciente) error
	Delete(id int) error
	Exists(codeValue string) bool
}

type TurnoInterface interface {
	Read(id int) (domain.Turno, error)
	Create(product domain.Turno) error
	Update(product domain.Turno) error
	Delete(id int) error
	Exists(codeValue string) bool
}
