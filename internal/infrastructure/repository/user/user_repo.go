package user_repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ArtoIi/To-Do-List-API/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(user *domain.User) error {

	query := `INSERT INTO user
			(name, email, hashed_password)
			VALUES (?,?,?)`

	_, err := r.db.Exec(query, user.Name, user.Email, user.HashedPassword)

	return err

}

func (r *UserRepository) GetEmail(email string) (*domain.User, error) {
	query := `SELECT id,name, email, hashed_password 
				FROM user
				Where email=? `

	user := &domain.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetId(id int) (*domain.User, error) {
	query := `SELECT id,name, email, hashed_password 
				FROM user
				Where id=? `

	user := &domain.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *domain.User) error {

	query := `UPDATE user SET name=?, email=?, hashed_password=? WHERE id=?`

	result, err := r.db.Exec(query, user.Name, user.Email, user.HashedPassword, user.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("nenhum registro foi atualizado")
	}

	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM user Where id=? `

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("nenhum post encontrado com o ID %d", id)
	}
	return nil
}
