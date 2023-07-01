package store

import (
	"database/sql"
	"fmt"

	"github.com/zoezn/Final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &SqlStore{
		DB: db,
	}
}

func (s *SqlStore) Read(id int) (domain.Dentista, error) {
	var productReturn domain.Dentista

	query := "SELECT * FROM dentista WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&productReturn.Id, &productReturn.Nombre, &productReturn.Apellido, &productReturn.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return productReturn, nil
}

func (s *SqlStore) Create(product domain.Dentista) error {
	query := "INSERT INTO dentista (Apellido, Nombre, Matricula) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(product.Apellido, product.Nombre, product.Matricula)
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

func (s *SqlStore) Update(product domain.Dentista) error {
	return nil
}

func (s *SqlStore) Delete(id int) error {
	d, err := s.Read(id)
	if err != nil {
		return err
	}
	query := "DELETE FROM dentista WHERE id = ?;"
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

	query := "SELECT * FROM dentista WHERE code_value = ?;"
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
