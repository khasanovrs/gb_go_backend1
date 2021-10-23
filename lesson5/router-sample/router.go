package router

import (
	"goback1/lesson5/reguser/api/handler"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Router struct {
	http.Handler
	hs *handler.Handlers
}

func NewRouter(hs *handler.Handlers) *Router {
	rh := &Router{
		hs: hs,
	}

	// r := mux.NewRouter()
	// r := chi.NewRouter()
	// r := gin.Default()

	// r.POST("/create", rh.CreateUser)

	// r.Post("/create", rh.CreateUser)
	// r.Get("/read/{id}", rh.ReadUser)
	// r.Delete("/delete/{id}", rh.DeleteUser)
	// r.Get("/search/{q}", rh.SearchUser)

	rh.Handler = r
	return rh
}

type User handler.User

// func (User) Bind(r *http.Request) error

// func (User) Render(w http.ResponseWriter, r *http.Request) error

func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	ru := User{}
	// if err := render.Bind(r, &ru); err != nil {
	// 	render.Render(w, r, ErrRender(err))
	// 	return
	// }

	u, err := rt.hs.CreateUser(r.Context(), handler.User(ru))
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Render(w, r, User(u))
}

func (rt *Router) ReadUser(w http.ResponseWriter, r *http.Request) {
	// sid := chi.URLParam(r, "id")

	uid, err := uuid.Parse(sid)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	u, err := rt.hs.ReadUser(r.Context(), uid)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Render(w, r, User(u))
}

func (rt *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// sid := chi.URLParam(r, "id")

	uid, err := uuid.Parse(sid)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	u, err := rt.hs.DeleteUser(r.Context(), uid)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Render(w, r, User(u))
}

func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {
	// q := chi.URLParam(r, "id")
}
