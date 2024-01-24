package routes

import (
	"database/sql"

	"theobourgeois.com/app/templates/pages/index"
	"theobourgeois.com/internal/router"
)

func InitRoutes(db *sql.DB) {
	// page routes
	router.CreateRoute("/", index.Index())

	// api routes
	// router.CreateApiRoute("/api/posts", http.MethodGet, posts.GetPosts(db))
}
