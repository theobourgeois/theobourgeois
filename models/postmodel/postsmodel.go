package postmodel

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	Id        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetById(id int, db *sql.DB) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"
	row := db.QueryRow(query, id)
	post := &Post{}
	var createdAt, updatedAt []byte
	err := row.Scan(&post.Id, &post.Title, &post.Body, &createdAt, &updatedAt)
	if err != nil {
		log.Println("Error getting post with query", query, err)
		return nil, err
	}
	post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		log.Println("Error parsing date", err)
		return nil, err
	}
	post.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		log.Println("Error parsing date", err)
		return nil, err
	}
	return post, nil
}

func GetAll(db *sql.DB) ([]*Post, error) {
	query := "SELECT * FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error getting posts with query", query, err)
		return nil, err
	}
	posts := []*Post{}
	for rows.Next() {
		post := &Post{}
		var createdAt, updatedAt []byte
		err := rows.Scan(&post.Id, &post.Title, &post.Body, &createdAt, &updatedAt)
		if err != nil {
			log.Println("Error getting post with query", query, err)
			continue
		}
		post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			log.Println("Error parsing date", err)
			continue
		}
		post.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if err != nil {
			log.Println("Error parsing date", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}
