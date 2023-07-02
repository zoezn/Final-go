package storeTurnos

import (
	"database/sql"
	"fmt"

	"github.com/zoezn/Final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) TurnoInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) Read(id int) (domain.Turno, error) {
	var pacienteReturn domain.Turno

	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&pacienteReturn.Id, &pacienteReturn.Dentista, &pacienteReturn.Paciente, &pacienteReturn.Descripcion, &pacienteReturn.Fecha, &pacienteReturn.Hora)
	if err != nil {
		return domain.Turno{}, err
	}
	return pacienteReturn, nil
}

func (s *SqlStore) Create(turno domain.Turno) error {
	query := "INSERT INTO turnos (dentista_id, paciente_id, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(turno.Dentista.Id, turno.Paciente.Id, turno.Fecha, turno.Hora, turno.Descripcion)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	turno.Id = int(lid)
	return nil
}

func (s *SqlStore) Update(id int, turno domain.Turno) error {
	query := "UPDATE pacientes SET `dentista_id` = ?, `paciente_id` = ?, `fecha` = ?, `hora` = ?, `descripcion` = ? WHERE `id` = ?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(turno.Dentista.Id, turno.Paciente.Id, turno.Fecha, turno.Hora, turno.Descripcion, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	turno.Id = int(lid)
	return nil
}

func (s *SqlStore) Delete(id int) error {
	d, err := s.Read(id)
	if err != nil {
		return err
	}
	query := "DELETE FROM turnos WHERE id = ?;"
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

	query := "SELECT * FROM turnos WHERE id = ?;"
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
