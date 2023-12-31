package turno

import (
	"github.com/zoezn/Final-go/internal/domain"
)

type TurnoService interface {
	GetByID(id int) (domain.Turno, error)
	GetByDNI(dni int) (domain.Turno, error)
	Create(p domain.Turno) (domain.Turno, error)
	CreateWithoutIds(p domain.Turno, dMatricula string, pDNI int) (domain.Turno, error)
	Delete(id int) error
	Update(id int, p domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) TurnoService {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}
func (s *service) GetByDNI(dni int) (domain.Turno, error) {
	t, err := s.r.GetByDNI(dni)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) Create(t domain.Turno) (domain.Turno, error) {
	turno, err := s.r.Create(t)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}
func (s *service) CreateWithoutIds(t domain.Turno, dMatricula string, pDNI int) (domain.Turno, error) {
	turno, err := s.r.CreateWithoutIds(t, dMatricula, pDNI)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}
func (s *service) Update(id int, u domain.Turno) (domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if t.Paciente != (domain.Paciente{}) {
		t.Paciente = u.Paciente
	}
	if u.Dentista != (domain.Dentista{}) {
		t.Dentista = u.Dentista
	}
	if u.Paciente != t.Paciente {
		t.Paciente = u.Paciente
	}
	if u.Dentista != t.Dentista {
		t.Dentista = u.Dentista
	}
	if u.Fecha != "" {
		t.Fecha = u.Fecha
	}

	if u.Hora != "" {
		t.Hora = u.Hora
	}

	if u.Descripcion != "" {
		t.Descripcion = u.Descripcion
	}

	tUpd, err := s.r.Update(id, t)
	if err != nil {
		return domain.Turno{}, err
	}
	return tUpd, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
