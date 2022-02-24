package database

import (
  // Why is Go like this :"(
  // Utilizing SIDE EFFECT
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

func ConnectDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

  return db, nil
}
