package sqlite

import (
	"database/sql"

	"github.com/kamil-budzik/gin-playground/internal/models"
)

type CharactersModel struct {
	DB *sql.DB
}

func (m *CharactersModel) All() ([]models.Character, error) {
	stmt := `SELECT id, name, role, bounty FROM characters ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	characters := []models.Character{}
	for rows.Next() {
		ch := models.Character{}

		err := rows.Scan(&ch.ID, &ch.Name, &ch.Role, &ch.Bounty)

		if err != nil {
			return nil, err
		}

		characters = append(characters, ch)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return characters, nil
}
