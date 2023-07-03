package store

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type DentistaInterface interface {
	Read(id int) (domain.Dentista, error)
	ReadByMatricula(matricula string) (domain.Dentista, error)
	Create(dentista domain.Dentista) error
	Update(id int, dentista domain.Dentista) error
	Delete(id int) error
	Exists(matricula string) bool
}
