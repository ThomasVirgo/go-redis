package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/redis/go-redis/v9"
)

func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	templates := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})
	router.HandleFunc("GET /{room_id}", func(w http.ResponseWriter, r *http.Request) {
		room_id := r.PathValue("room_id")
		val, err := rdb.Get(ctx, room_id).Result()
		if err != nil {
			w.Write([]byte("error reading from redis store"))
		}
		w.Write([]byte(val))
	})
	router.HandleFunc("POST /{room_id}", func(w http.ResponseWriter, r *http.Request) {
		room_id := r.PathValue("room_id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error reading body"))
		}
		err = rdb.Set(ctx, room_id, body, 0).Err()
		if err != nil {
			w.Write([]byte("error setting data in redis for room"))
		}
		w.Write([]byte(fmt.Sprintf("successfully updated room: %s", room_id)))
	})

	http.ListenAndServe(":9000", router)
}
