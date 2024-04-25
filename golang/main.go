package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	Id   string
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/bench")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var user User
		user.Id = uuid.NewString()

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Erro ao decodificar o JSON: %v", err)
			return
		}

		_, err = db.Exec("INSERT INTO users (id, name) values (?, ?)", user.Id, user.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Erro ao salvar os dados no banco de dados: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	fmt.Println("Servidor iniciado em :3000")
	http.ListenAndServe(":3000", r)
}
