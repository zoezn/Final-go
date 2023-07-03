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
	var turnoReturn domain.Turno
	var dentistaId int
	var pacienteId int

	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&turnoReturn.Id, &dentistaId, &pacienteId, &turnoReturn.Fecha, &turnoReturn.Hora, &turnoReturn.Descripcion)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	dQuery := "SELECT * FROM dentistas WHERE id = ?;"
	dRow := s.DB.QueryRow(dQuery, dentistaId)
	err = dRow.Scan(&turnoReturn.Dentista.Id, &turnoReturn.Dentista.Nombre, &turnoReturn.Dentista.Apellido, &turnoReturn.Dentista.Matricula)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	pQuery := "SELECT * FROM pacientes WHERE id = ?;"
	pRow := s.DB.QueryRow(pQuery, pacienteId)
	err = pRow.Scan(&turnoReturn.Paciente.Id, &turnoReturn.Paciente.Nombre, &turnoReturn.Paciente.Apellido, &turnoReturn.Paciente.Domicilio, &turnoReturn.Paciente.DNI, &turnoReturn.Paciente.Alta)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	return turnoReturn, nil
}
func (s *SqlStore) ReadByDni(dni int) (domain.Turno, error) {
	var turnoReturn domain.Turno
	var dentistaId int
	var pacienteId int

	query := "SELECT * FROM turnos t JOIN pacientes p ON t.paciente_id = p.id WHERE p.dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&turnoReturn.Id, &dentistaId, &pacienteId, &turnoReturn.Fecha, &turnoReturn.Hora, &turnoReturn.Descripcion)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	dQuery := "SELECT * FROM dentistas WHERE id = ?;"
	dRow := s.DB.QueryRow(dQuery, dentistaId)
	err = dRow.Scan(&turnoReturn.Dentista.Id, &turnoReturn.Dentista.Nombre, &turnoReturn.Dentista.Apellido, &turnoReturn.Dentista.Matricula)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	pQuery := "SELECT * FROM pacientes WHERE id = ?;"
	pRow := s.DB.QueryRow(pQuery, pacienteId)
	err = pRow.Scan(&turnoReturn.Paciente.Id, &turnoReturn.Paciente.Nombre, &turnoReturn.Paciente.Apellido, &turnoReturn.Paciente.Domicilio, &turnoReturn.Paciente.DNI, &turnoReturn.Paciente.Alta)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}

	return turnoReturn, nil
}

func (s *SqlStore) Create(turno domain.Turno) error {
	query := "INSERT INTO turnos (dentista_id, paciente_id, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := stmt.Exec(turno.Dentista.Id, turno.Paciente.Id, turno.Fecha, turno.Hora, turno.Descripcion)
	if err != nil {
		fmt.Println(err)
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

func (s *SqlStore) CreateWithoutIds(turno domain.Turno, dentistaId int, pacienteId int) error {

	query := "INSERT INTO turnos (dentista_id, paciente_id, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := stmt.Exec(dentistaId, pacienteId, turno.Fecha, turno.Hora, turno.Descripcion)
	if err != nil {
		fmt.Println(err)
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
	query := "UPDATE turnos SET `dentista_id` = ?, `paciente_id` = ?, `fecha` = ?, `hora` = ?, `descripcion` = ? WHERE `id` = ?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(turno.Dentista.Id, turno.Paciente.Id, turno.Fecha, turno.Hora, turno.Descripcion, id)
	fmt.Println(err)
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
