package store

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type StoreInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Dentista, error)
	// Create agrega un nuevo producto
	Create(product domain.Dentista) error
	// Update actualiza un producto
	Update(product domain.Dentista) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}

type PacienteInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Paciente, error)
	// Create agrega un nuevo producto
	Create(product domain.Paciente) error
	// Update actualiza un producto
	Update(product domain.Paciente) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}

type TurnoInterface interface {
	// Read devuelve un turno por su id
	Read(id int) (domain.Turno, error)
	// Create agrega un nuevo turno
	Create(product domain.Turno) error
	// Update actualiza un turno
	Update(product domain.Turno) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
