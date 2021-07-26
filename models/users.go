package users

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/pgxpool"
)

var ( 
	id int
	name string
)
}

func main() {
	db, err := sql.Open("psql", 
		"postgres:postgres@tcp(127.0.0.1:5302)/go_blog")
	if err != nil {
		log.Fatal(err)
	}

	rows,err := db.Query("select id, naem from users where id = ?", 1)
	if err !- nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}