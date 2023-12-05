package stdhttp

import (
	// "encoding/json"
	// "fmt"
	"context"
	"httpserver/gates/psg"
	// "httpserver/models/dto"
	// "httpserver/pkg"
	// "io"
	// "log"
	"net/http"

	"github.com/pkg/errors"
	// "github.com/gorilla/mux"
)

type Controller struct {
	srv http.Server
	db  *psg.Psg
}

func NewController(addr string, postgres *psg.Psg) (hs *Controller) {
	hs = new(Controller)
	hs.srv = http.Server{}
	mux := http.NewServeMux()

	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) /// наверно также как обычные ошибки нужно обрабатывать
			return
		}
		hs.RecordCreateHandler(w, r)
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		hs.RecordGetHandler(w, r)
	})
	mux.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		hs.RecordUpdateHandler(w, r)
	})
	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		hs.RecordDeleteByPhoneHandler(w, r)
	})
	hs.srv.Handler = mux
	hs.srv.Addr = addr
	hs.db = postgres
	return hs
}

func (psg *Psg) NoteSave(name, lastName, note string) (err error) {
	defer func() { err = errors.Wrap(err, "postgres NoteSave()") }()

	query := "INSERT INTO notes (name, last_name, note) VALUES ($1, $2, $3)"

	_, err = psg.conn.Exec(context.Background(), query, name, lastName, note)
	if err != nil {
		err = errors.Wrap(err, "psg.conn.Exec(context.Background(), query, name, lastName, note)")
		return
	}
	return
}

func (psg *Psg) NoteRead(name, lastName, note string) (err error) {
	// TODO
}

func (psg *Psg) NoteDelete(name, lastName, note string) (err error) {
	// TODO
}
