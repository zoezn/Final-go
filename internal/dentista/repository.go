package dentista

import (
	"errors"

	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	storage store.DentistaInterface
}

func NewRepository(storage store.DentistaInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}
	return dentista, nil

}

func (r *repository) Create(p domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(p.Matricula) {
		return domain.Dentista{}, errors.New("Matricula already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Dentista{}, errors.New("Error creating dentista")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(p.Matricula) {
		return domain.Dentista{}, errors.New("Matricula already exists")
	}
	err := r.storage.Update(id, p)
	if err != nil {
		return domain.Dentista{}, errors.New("Error updating dentista")
	}
	return p, nil
}
