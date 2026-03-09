package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(dns string) (domain.ToDoRepository, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLRepository{db: db}, nil
}

func (r *MySQLRepository) Register(DTO domain.CreateUserDTO) error {

	query := `INSERT INTO user
			(name, email, hashed_password)
			VALUES (?,?,?)`

	hashedpassword, err := utils.HashedPassword(DTO.Password)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, DTO.Name, DTO.Email, hashedpassword)
	return err
}
