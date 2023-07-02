package storeTurnos

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type TurnoInterface interface {
	Read(id int) (domain.Turno, error)
	Create(turno domain.Turno) error
	Update(id int, turno domain.Turno) error
	Delete(id int) error
	Exists(id int) bool
}
