package storePaciente

import (
	"database/sql"
	"fmt"

	"github.com/zoezn/Final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) PacienteInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) Read(id int) (domain.Paciente, error) {
	var pacienteReturn domain.Paciente

	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&pacienteReturn.Id, &pacienteReturn.Nombre, &pacienteReturn.Apellido, &pacienteReturn.Domicilio, &pacienteReturn.DNI, &pacienteReturn.Alta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return pacienteReturn, nil
}
func (s *SqlStore) ReadByDNI(dni int) (domain.Paciente, error) {
	var pacienteReturn domain.Paciente

	query := "SELECT * FROM pacientes WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&pacienteReturn.Id, &pacienteReturn.Nombre, &pacienteReturn.Apellido, &pacienteReturn.Domicilio, &pacienteReturn.DNI, &pacienteReturn.Alta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return pacienteReturn, nil
}

func (s *SqlStore) Create(paciente domain.Paciente) error {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.Alta)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	paciente.Id = int(lid)
	fmt.Println(paciente)
	return nil
}

func (s *SqlStore) Update(id int, paciente domain.Paciente) error {
	query := "UPDATE pacientes SET `nombre` = ?, `apellido` = ?, `domicilio` = ?, `dni` = ?, `alta` = ? WHERE `id` = ?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.Alta, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	paciente.Id = int(lid)
	fmt.Println(paciente)
	return nil
}

func (s *SqlStore) Delete(id int) error {
	d, err := s.Read(id)
	if err != nil {
		return err
	}
	query := "DELETE FROM pacientes WHERE id = ?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(d)
	return nil
}

func (s *SqlStore) Exists(id int) bool {
	var exist bool
	var dbId int

	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dbId)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}
func (s *SqlStore) ExistsByDni(dni int) bool {
	var exist bool
	var dbId int

	query := "SELECT * FROM pacientes WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&dbId)
	if err != nil {
		return exist
	}

	if dni >= 0 {
		exist = true
	}

	return exist

}
