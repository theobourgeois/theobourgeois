package posts

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"theobourgeois.com/app/models/postmodel"
	tpost "theobourgeois.com/app/templates/components/postcomps"
	"theobourgeois.com/internal/router"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func PostPosts(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		_, err := db.Exec("INSERT INTO posts (title, body) VALUES (?, ?)", r.FormValue("name"), r.FormValue("body"))
		if err != nil {
			http.Error(w, "Error inserting post", http.StatusInternalServerError)
			log.Println("Error inserting post", err)
			return templ.NopComponent
		}

		posts, err := postmodel.GetAll(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return templ.NopComponent
		}
		return tpost.Posts(posts)
	}
}

func GetPosts(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		//time.Sleep(500 * time.Millisecond)
		posts, err := postmodel.GetAll(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return templ.NopComponent
		}
		return tpost.Posts(posts)
	}
}

func GetPost(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Error parsing id", http.StatusInternalServerError)
			log.Println("Error parsing id", err)
			return templ.NopComponent
		}

		post, err := postmodel.GetById(id, db)
		if err != nil {
			http.Error(w, "Error getting post", http.StatusInternalServerError)
			log.Println("Error getting post", err)
			return templ.NopComponent
		}
		log.Println("Getting post with id", id)
		return tpost.Post(post)
	}
}

func DeletePost(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		id := mux.Vars(r)["id"]

		_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error deleting post", http.StatusInternalServerError)
			log.Println("Error deleting post", err)
		}

		posts, err := postmodel.GetAll(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return templ.NopComponent
		}
		return tpost.Posts(posts)
	}
}

// ToggleEditPost toggles the edit state of a post
// If the query param 'editing' is true, it will return the editpost component
// Otherwise, it will return the post component
func ToggleEditPost(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		editing := r.URL.Query().Get("editing") == "true"

		if err != nil {
			http.Error(w, "Error parsing id", http.StatusInternalServerError)
			log.Println("Error parsing id", err)
			return templ.NopComponent
		}

		post, err := postmodel.GetById(id, db)
		if err != nil {
			http.Error(w, "Error getting post", http.StatusInternalServerError)
			log.Println("Error getting post", err)
			return templ.NopComponent
		}

		if editing {
			log.Println("Editing post with id", id)
			return tpost.EditPost(post)
		}

		return tpost.Post(post)
	}
}

func UpdatePost(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Error parsing id", http.StatusInternalServerError)
			log.Println("Error parsing id", err)
			return templ.NopComponent
		}

		_, err = db.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", r.FormValue("name"), r.FormValue("body"), id)

		if err != nil {
			http.Error(w, "Error updating post", http.StatusInternalServerError)
			log.Println("Error updating post", err)
			return templ.NopComponent
		}

		post, err := postmodel.GetById(id, db)

		if err != nil {
			http.Error(w, "Error getting post", http.StatusInternalServerError)
			log.Println("Error getting post", err)
			return templ.NopComponent
		}

		return tpost.Post(post)
	}
}
