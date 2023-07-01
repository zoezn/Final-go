package paciente

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
	var productReturn domain.Paciente

	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&productReturn.Id, &productReturn.Nombre, &productReturn.Apellido, &productReturn.Id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return productReturn, nil
}

func (s *SqlStore) Create(product domain.Paciente) error {
	query := "INSERT INTO pacientes (Apellido, Nombre, Id) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(product.Apellido, product.Nombre, product.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	lid, _ := res.LastInsertId()
	product.Id = int(lid)
	fmt.Println(product)
	return nil
}

func (s *SqlStore) Update(product domain.Paciente) error {
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

func (s *SqlStore) Exists(codeValue string) bool {
	var exist bool
	var id int

	query := "SELECT * FROM pacientes WHERE code_value = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}
