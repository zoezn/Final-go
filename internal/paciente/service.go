package paciente

import (
	"strconv"

	"github.com/zoezn/Final-go/internal/domain"
)

type PacienteService interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) PacienteService {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Domicilio != "" {
		p.Domicilio = u.Domicilio
	}

	if strconv.Itoa(u.DNI) != "" {
		p.DNI = u.DNI
	}

	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
