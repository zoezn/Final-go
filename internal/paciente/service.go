package paciente

import (
	"strconv"

	"github.com/zoezn/Final-go/internal/domain"
)

type PacienteService interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

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
