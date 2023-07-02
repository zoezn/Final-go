package dentista

import (
	"errors"

	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/pkg/store"
)

type Repository interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo producto
	Create(p domain.Dentista) (domain.Dentista, error)
	// Update actualiza un producto
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un producto
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Dentista, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}
	return product, nil

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
