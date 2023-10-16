package workdb

import (
	"database/sql"
	"fmt"
	comments "practiceDB/requests/coment"
	posts "practiceDB/requests/post"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DataConfig struct {
	UserName string
	Password string
	DbName   string
}

// підключення до БД
func (d *DataConfig) Conection() (*sql.DB, error) {
	var metaData string = fmt.Sprintf("%s:%s@/%s", d.UserName, d.Password, d.DbName)
	db, err := sql.Open("mysql", metaData)

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

// запис постів до БД
func AddPost(data posts.Post, wg *sync.WaitGroup, db *sql.DB) error {

	defer wg.Done()

	query := "INSERT INTO posts (iduser, id, title, body) VALUES (?, ?, ?, ?)"

	_, err := db.Exec(query, data.UserId, data.PostId, data.Title, data.Body)

	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Post added successfully.")
	return nil
}

// запис коментарів до БД
func AddComment(data comments.Coments, wg *sync.WaitGroup, db *sql.DB) error {
	defer wg.Done()

	query := "INSERT INTO coments (postid, id, name, email, body) VALUES (?, ?, ?, ?, ?)"

	_, err := db.Exec(query, data.PostId, data.ComentsId, data.ComentsName, data.UserEmail, data.ComentsBody)

	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Post added successfully.")
	return nil
}
