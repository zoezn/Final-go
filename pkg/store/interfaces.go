package store

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type DentistaInterface interface {
	Read(id int) (domain.Dentista, error)
	Create(dentista domain.Dentista) error
	Update(id int, dentista domain.Dentista) error
	Delete(id int) error
	Exists(codeValue string) bool
}

type TurnoInterface interface {
	Read(id int) (domain.Turno, error)
	Create(turno domain.Turno) error
	Update(turno domain.Turno) error
	Delete(id int) error
	Exists(codeValue string) bool
}
