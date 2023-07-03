package storeTurnos

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type TurnoInterface interface {
	Read(id int) (domain.Turno, error)
	ReadByDni(dni int) (domain.Turno, error)
	Create(turno domain.Turno) error
	CreateWithoutIds(turno domain.Turno, dentistaId int, pacienteId int) error
	Update(id int, turno domain.Turno) error
	Delete(id int) error
	Exists(id int) bool
}
