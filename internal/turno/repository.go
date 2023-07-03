package turno

import (
	"errors"

	"github.com/zoezn/Final-go/internal/domain"
	"github.com/zoezn/Final-go/pkg/store"
	"github.com/zoezn/Final-go/pkg/storePaciente"
	"github.com/zoezn/Final-go/pkg/storeTurnos"
)

type Repository interface {
	GetByID(id int) (domain.Turno, error)
	GetByDNI(dni int) (domain.Turno, error)
	Create(t domain.Turno) (domain.Turno, error)
	CreateWithoutIds(t domain.Turno, dMatricula string, pDNI int) (domain.Turno, error)
	Update(id int, t domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

type repository struct {
	storage          storeTurnos.TurnoInterface
	storagePacientes storePaciente.PacienteInterface
	storageDentistas store.DentistaInterface
}

func NewRepository(storage storeTurnos.TurnoInterface, storagePacientes storePaciente.PacienteInterface,
	storageDentistas store.DentistaInterface) Repository {
	return &repository{storage, storagePacientes, storageDentistas}
}

func (r *repository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, errors.New("Turno not found")
	}
	return turno, nil
}

func (r *repository) GetByDNI(dni int) (domain.Turno, error) {
	turno, err := r.storage.Read(dni)
	if err != nil {
		return domain.Turno{}, errors.New("Turno not found")
	}
	return turno, nil
}

func (r *repository) Create(t domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(t.Id) {
		return domain.Turno{}, errors.New("id already exists")
	}
	err := r.storage.Create(t)
	if err != nil {
		return domain.Turno{}, errors.New("error creating Turno")
	}
	return t, nil
}
func (r *repository) CreateWithoutIds(t domain.Turno, dMatricula string, pDNI int) (domain.Turno, error) {
	if r.storageDentistas.Exists(dMatricula) {
		return domain.Turno{}, errors.New("Dentista doesn't exist")
	}

	dByMatricula, err := r.storageDentistas.ReadByMatricula(dMatricula)
	if err != nil {
		return domain.Turno{}, errors.New("Error creating Turno: no dentista with that Matricula")
	}

	if r.storagePacientes.Exists(pDNI) {
		return domain.Turno{}, errors.New("Paciente doesn't exist")
	}
	pByDNI, err := r.storagePacientes.ReadByDNI(pDNI)
	if err != nil {
		return domain.Turno{}, errors.New("Error creating Turno: no paciente with that DNI")
	}
	if r.storage.Exists(t.Id) {
		return domain.Turno{}, errors.New("id already exists")
	}
	err = r.storage.CreateWithoutIds(t, dByMatricula.Id, pByDNI.Id)
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
	if r.storage.Exists(t.Id) {
		return domain.Turno{}, errors.New("id already exists")
	}
	err := r.storage.Update(id, t)
	if err != nil {
		return domain.Turno{}, errors.New("error updating Turno")
	}
	return t, nil
}
