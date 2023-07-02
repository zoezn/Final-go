package turno

import (
	"errors"
	"strconv"

	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Turno, error)
	Create(t domain.Turno) (domain.Turno, error)
	Update(id int, t domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

type repository struct {
	storage store.TurnoInterface
}

func NewRepository(storage store.TurnoInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, errors.New("Turno not found")
	}
	return turno, nil

}

func (r *repository) Create(t domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(strconv.Itoa(t.Id)) {
		return domain.Turno{}, errors.New("dni already exists")
	}
	err := r.storage.Create(t)
	if err != nil {
		return domain.Turno{}, errors.New("error creating Turno")
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, t domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(strconv.Itoa(t.Id)) {
		return domain.Turno{}, errors.New("dni already exists")
	}
	err := r.storage.Update(t)
	if err != nil {
		return domain.Turno{}, errors.New("error updating Turno")
	}
	return t, nil
}
