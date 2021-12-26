package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type info struct {
	Name      string
	Age       int16
	Education string
	image     string
}

func infoByName(name string, db *sql.DB) ([]info, error) {
	var Info []info
	rows, err := db.Query("SELECT * FROM info WHERE Name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb info
		if err := rows.Scan(&alb.Name, &alb.Age, &alb.Education, &alb.image); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		Info = append(Info, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return Info, nil
}
func addName(In info, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO `info`(`Name`, `Age`, `Education`, `Image`) VALUES (?,?,?,?)", In.Name, In.Age, In.Education, In.image)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
func main() {
	db, err := sql.Open("mysql", "subtawee:1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	Info, err := infoByName("Subtawee nganrungruang", db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", Info)
	albID, err := addName(info{
		Name:      "nitirod",
		Age:       15,
		Education: "kmutt",
		image:     "address",
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}
