package todo_repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ArtoIi/To-Do-List-API/internal/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Post(todo *domain.ToDo) error {

	query := `INSERT INTO todo
			(user_id, title, description, created_at, updated_at)
			VALUES (?,?,?,?,?)`

	_, err := r.db.Exec(query, todo.UserID, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt)

	return err

}

func (r *TodoRepository) GetId(id int) (*domain.ToDo, error) {
	query := `SELECT id,user_id, title, description, created_at, updated_at 
				FROM todo
				Where id=? `

	todo := &domain.ToDo{}
	err := r.db.QueryRow(query, id).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) GetUserId(userID int, limit, offset int) ([]*domain.ToDo, int, error) {

	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM todo WHERE user_id = ?", userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id,user_id, title, description, created_at, updated_at 
				FROM todo
				Where user_id=?
				LIMIT ? 
				OFFSET ?`

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var todos []*domain.ToDo
	for rows.Next() {
		todo := &domain.ToDo{}

		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Description,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		todos = append(todos, todo)
	}

	return todos, total, nil
}

func (r *TodoRepository) Update(todo *domain.ToDo) (*domain.ToDo, error) {

	query := `UPDATE todo SET title=?, description=?, updated_at=? WHERE id=?`

	result, err := r.db.Exec(query, todo.Title, todo.Description, todo.UpdatedAt, todo.ID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("nenhum registro foi atualizado")
	}

	return todo, nil
}

func (r *TodoRepository) Delete(id int) error {
	query := `DELETE FROM todo Where id=? `

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
