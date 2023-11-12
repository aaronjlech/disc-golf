package repository

import (
	"database/sql"
	"fmt"
)

// saveOrUpdateManufacturer inserts or updates data in the database
func CreateNewDisc(db *sql.DB, name, href string) error {
	// Use an UPSERT (INSERT ON CONFLICT UPDATE) statement
	fmt.Println("saving", href, name)
	query := `
		INSERT INTO public.discs (name, href)
		VALUES ($1, $2)
		ON CONFLICT (name) DO UPDATE
		SET href = EXCLUDED.href;
	`

	_, err := db.Exec(query, name, href)
	fmt.Println("ERRR %s", err)
	return err
}
