package creatdb

import (
	"new/bcryptp"
	"new/database"

	"github.com/mattn/go-sqlite3"
)

func Creatdb() error {
	db := database.GetDb()

	// Create table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL
	);`)
	if err != nil {
		return err
	}

	// Vérifier si l'admin existe déjà
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = 'Admin'").Scan(&count)
	if err != nil {
		return err
	}

	// Si aucun utilisateur n'existe, créer les utilisateurs par défaut
	if count == 0 {
		// Hash passwords before insertion
		adminPass, err := bcryptp.HashPassword("123")
		if err != nil {
			return err
		}
		vendeur1Pass, err := bcryptp.HashPassword("123")
		if err != nil {
			return err
		}
		vendeur2Pass, err := bcryptp.HashPassword("123")
		if err != nil {
			return err
		}

		// Insert users with hashed passwords
		_, err = db.Exec(`
		INSERT INTO users (name, password, role) VALUES
		('Admin', ?, 'admin'),
		('vendeur1', ?, 'vendeur'),
		('vendeur2', ?, 'vendeur');
		`, adminPass, vendeur1Pass, vendeur2Pass)

		if err != nil {
			// Si l'erreur est due à un conflit de clé unique, on ignore
			if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
				return nil
			}
			return err
		}
	}

	return nil
}
