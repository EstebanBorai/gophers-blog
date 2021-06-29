package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Api struct {
	router     http.Handler
	dataSource DataSource
}

type Server interface {
	Router() http.Handler
	Serve()
}

func NewServer(dataSource DataSource) Server {
	api := new(Api)
	router := mux.NewRouter()

	router.HandleFunc("/", api.showPosts).Methods("GET")
	router.HandleFunc("/{id}", api.findPost).Methods("GET")
	router.HandleFunc("/", api.createPost).Methods("POST")
	router.HandleFunc("/{id}", api.updatePost).Methods("PUT")
	router.HandleFunc("/{id}", api.deletePost).Methods("DELETE")

	api.router = router
	api.dataSource = dataSource

	return api
}

func (a *Api) Router() http.Handler {
	return a.router
}

func (a *Api) Serve() {
	log.Println("Binding server to http://127.0.0.1:8080")

	if err := http.ListenAndServe(":8080", a.Router()); err != nil {
		log.Fatal(err)
	}
}

func (a *Api) respondErr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(err.Error()))
}

func (a *Api) showPosts(w http.ResponseWriter, r *http.Request) {
	if posts, err := a.dataSource.ShowPosts(); err != nil {
		a.respondErr(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

func (a *Api) findPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		a.respondErr(w, err)
		return
	}

	if posts, err := a.dataSource.FindPost(id); err != nil {
		a.respondErr(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

func (a *Api) createPost(w http.ResponseWriter, r *http.Request) {
	var dto PostDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		a.respondErr(w, err)
	}

	if post, err := a.dataSource.InsertPost(&dto); err != nil {
		a.respondErr(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func (a *Api) updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		a.respondErr(w, err)
		return
	}

	var dto PostDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		a.respondErr(w, err)
	}

	if post, err := a.dataSource.UpdatePost(id, &dto); err != nil {
		a.respondErr(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func (a *Api) deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		a.respondErr(w, err)
		return
	}

	if deletedPost, err := a.dataSource.DeletePost(id); err != nil {
		a.respondErr(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deletedPost)
	}
}
