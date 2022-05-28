/*
 * vidhublog: Vidhu Kant's Blog
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
  "os"
)

// ConnectDB opens a connection to the database
// TODO: properly handle all errors
func ConnectDB() *sql.DB {
  // load ENV
  err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

  // read ENV
  username := os.Getenv("DB_USERNAME")
  password := os.Getenv("DB_PASSWORD")
  database := os.Getenv("DB_DATABASE")

  // connect to database
  db, err := sql.Open("mysql", username + ":" + password + "@/" + database)
	if err != nil {
		panic(err.Error())
	}

	return db
}
