package database

import (
	"database/sql"
	"new/bcryptp"
)

func GetLogin(name, password string) (bool, string, error) {
	var hashedPassword, role string
	err := db.QueryRow("SELECT password, role FROM users WHERE name = ?", name).Scan(&hashedPassword, &role)
	
	if err == sql.ErrNoRows {
		return false, "", nil
	}
	if err != nil {
		return false, "", err
	}

	if bcryptp.CheckPasswordHash(password, hashedPassword) {
		return true, role, nil
	}

	return false, "", nil
}
