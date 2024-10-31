package repo

import (
	"fortress/model"
	"github.com/jmoiron/sqlx"
)

func AddUser(connection *sqlx.DB, model *model.User) (*model.User, error) {
	stmt := `INSERT INTO users VALUES(NULL, ?, ?, ?, ?, ?, datetime(CURRENT_TIMESTAMP, 'utc'))`
	res, err := connection.Exec(
		stmt,
		model.Username,
		model.NameFirst,
		model.NameLast,
		model.Email,
		model.Password,
	)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	model.Id = int(id)
	return model, err
}
