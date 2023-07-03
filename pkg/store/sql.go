package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/zoezn/Final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) DentistaInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) Read(id int) (domain.Dentista, error) {
	var dentistaReturn domain.Dentista

	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistaReturn.Id, &dentistaReturn.Nombre, &dentistaReturn.Apellido, &dentistaReturn.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentistaReturn, nil
}
func (s *SqlStore) ReadByMatricula(matricula string) (domain.Dentista, error) {
	var dentistaReturn domain.Dentista

	query := "SELECT * FROM dentistas WHERE matricula = ?;"
	row := s.DB.QueryRow(query, matricula)
	err := row.Scan(&dentistaReturn.Id, &dentistaReturn.Nombre, &dentistaReturn.Apellido, &dentistaReturn.Matricula)
	if err != nil {
		return domain.Dentista{}, errors.New(dentistaReturn.Matricula)
	}
	return dentistaReturn, nil
}

func (s *SqlStore) Create(dentista domain.Dentista) error {
	query := "INSERT INTO dentistas (Apellido, Nombre, Matricula) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	dentista.Id = int(lid)
	fmt.Println(dentista)
	return nil
}

func (s *SqlStore) Update(id int, dentista domain.Dentista) error {
	query := "UPDATE dentistas SET `apellido` = ?,`nombre` = ?,`matricula` = ? WHERE `id` = ?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	dentista.Id = int(lid)
	fmt.Println(dentista)
	return nil
}

func (s *SqlStore) Delete(id int) error {
	d, err := s.Read(id)
	if err != nil {
		return err
	}
	query := "DELETE FROM dentistas WHERE id = ?;"
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

func (s *SqlStore) Exists(matricula string) bool {
	var exist bool

	query := "SELECT * FROM dentistas WHERE matricula = ?;"
	row := s.DB.QueryRow(query, matricula)
	err := row.Scan(&matricula)
	if err != nil {
		return exist
	}

	if matricula != "" {
		exist = true
	}

	return exist
}
